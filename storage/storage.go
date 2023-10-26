package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/Filecoin-Titan/titan/api"
	"github.com/Filecoin-Titan/titan/api/client"
	"github.com/Filecoin-Titan/titan/api/types"
	cliutil "github.com/Filecoin-Titan/titan/cli/util"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/ipfs/go-cid"
	"github.com/zscboy/titan-game-sdk/storage/memfile"
)

type FileType string

const (
	FileTypeFile   FileType = "file"
	FileTypeFolder FileType = "folder"
)

type storageClose func()
type progressFunc func(doneSize int64, totalSize int64)

type Storage interface {
	// UploadFile the file path can be the absolute path of a single file or a directory
	UploadFile(ctx context.Context, filePath string, progress progressFunc) (cid.Cid, error)
	UploadStream(ctx context.Context, r io.Reader, progress progressFunc) (cid.Cid, error)
	DeleteFile(ctx context.Context, rootCID string) error
}

type storage struct {
	schedulerAPI api.Scheduler
}

func NewStorage(locatorURL, apiKey string) (Storage, storageClose, error) {
	udpPacketConn, err := net.ListenPacket("udp", ":0")
	if err != nil {
		return nil, nil, fmt.Errorf("ListenPacket %w", err)
	}

	// use http3 client
	httpClient, err := cliutil.NewHTTP3Client(udpPacketConn, true, "")
	if err != nil {
		return nil, nil, fmt.Errorf("NewHTTP3Client %w", err)
	}

	locatorAPI, _, err := client.NewLocator(context.TODO(), locatorURL, nil, jsonrpc.WithHTTPClient(httpClient))
	if err != nil {
		return nil, nil, fmt.Errorf("NewLocator %w", err)
	}

	schedulerURL, err := locatorAPI.GetSchedulerWithAPIKey(context.Background(), apiKey)
	if err != nil {
		return nil, nil, fmt.Errorf("GetSchedulerWithAPIKey %w", err)
	}

	headers := http.Header{}
	headers.Add("Authorization", "Bearer "+apiKey)

	schedulerAPI, apiClose, err := client.NewScheduler(context.TODO(), schedulerURL, headers, jsonrpc.WithHTTPClient(httpClient))
	if err != nil {
		return nil, nil, fmt.Errorf("NewScheduler %w", err)
	}

	close := func() {
		apiClose()
		udpPacketConn.Close()
	}
	return &storage{schedulerAPI: schedulerAPI}, close, nil
}

// The file path can be a single file or a directory
func (s *storage) UploadFile(ctx context.Context, filePath string, progress progressFunc) (cid.Cid, error) {
	// delete template file if exist
	fileName := filepath.Base(filePath)
	tempFile := path.Join(os.TempDir(), fileName)
	if _, err := os.Stat(tempFile); err == nil {
		os.Remove(tempFile)
	}

	root, err := createCar(filePath, tempFile)
	if err != nil {
		return cid.Cid{}, err
	}

	carFile, err := os.Open(tempFile)
	if err != nil {
		return cid.Cid{}, err
	}

	defer func() {
		if err = carFile.Close(); err != nil {
			fmt.Println("close car file error ", err.Error())
		}

		if err = os.Remove(tempFile); err != nil {
			fmt.Println("delete temporary car file error ", err.Error())
		}
	}()

	fileInfo, err := carFile.Stat()
	if err != nil {
		return cid.Cid{}, err
	}

	fileType, err := getFileType(filePath)
	if err != nil {
		return cid.Cid{}, err
	}

	assetProperty := &types.AssetProperty{AssetCID: root.String(), AssetName: fileName, AssetSize: fileInfo.Size(), AssetType: fileType}

	rsp, err := s.schedulerAPI.CreateUserAsset(ctx, assetProperty)
	if err != nil {
		return cid.Cid{}, fmt.Errorf("CreateUserAsset error %w", err)
	}

	if rsp.AlreadyExists {
		return cid.Cid{}, fmt.Errorf("asset %s already exist", root.String())
	}

	err = s.uploadFileWithForm(ctx, carFile, fileName, rsp.UploadURL, rsp.Token, progress)
	if err != nil {
		return cid.Cid{}, fmt.Errorf("uploadFileWithForm error %w", err)
	}

	return root, nil
}

func (s *storage) uploadFileWithForm(ctx context.Context, r io.Reader, name, uploadURL, token string, progress progressFunc) error {
	// Create a new multipart form body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a new form field for the file
	fileField, err := writer.CreateFormFile("file", name)
	if err != nil {
		return err
	}

	// Copy the file data to the form field
	_, err = io.Copy(fileField, r)
	if err != nil {
		return err
	}

	// Close the multipart form
	err = writer.Close()
	if err != nil {
		return err
	}

	totalSize := body.Len()
	dongSize := int64(0)
	pr := &ProgressReader{body, func(r int64) {
		if r > 0 {
			dongSize += r
			progress(dongSize, int64(totalSize))
		}
	}}

	// Create a new HTTP request with the form data
	request, err := http.NewRequest("POST", uploadURL, pr)
	if err != nil {
		return fmt.Errorf("new request error %s", err.Error())
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", "Bearer "+token)
	request = request.WithContext(ctx)

	// Create an HTTP client and send the request
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("do error %s", err.Error())
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	type result struct {
		Code int    `json:"code"`
		Err  int    `json:"err"`
		Msg  string `json:"msg"`
	}

	ret := result{}
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if ret.Code != 0 {
		return fmt.Errorf(ret.Msg)
	}

	return nil
}

func getFileType(filePath string) (string, error) {
	fileType := FileTypeFile
	if fileInfo, err := os.Stat(filePath); err != nil {
		return "", err
	} else if fileInfo.IsDir() {
		fileType = FileTypeFolder
	}

	return string(fileType), nil
}

func (s *storage) DeleteFile(ctx context.Context, rootCID string) error {
	return s.schedulerAPI.DeleteUserAsset(ctx, rootCID)
}

func (s *storage) UploadStream(ctx context.Context, r io.Reader, progress progressFunc) (cid.Cid, error) {
	memFile := memfile.New([]byte{})
	root, err := createCarStream(r, memFile)
	if err != nil {
		return cid.Cid{}, err
	}

	assetProperty := &types.AssetProperty{AssetCID: root.String(), AssetName: root.String(), AssetSize: int64(len(memFile.Bytes())), AssetType: string(FileTypeFile)}

	rsp, err := s.schedulerAPI.CreateUserAsset(ctx, assetProperty)
	if err != nil {
		return cid.Cid{}, fmt.Errorf("CreateUserAsset error %w", err)
	}

	if rsp.AlreadyExists {
		return cid.Cid{}, fmt.Errorf("asset %s already exist", root.String())
	}

	err = s.uploadFileWithForm(ctx, memFile, root.String(), rsp.UploadURL, rsp.Token, progress)
	if err != nil {
		return cid.Cid{}, fmt.Errorf("uploadFileWithForm error %w", err)
	}

	return root, nil
}
