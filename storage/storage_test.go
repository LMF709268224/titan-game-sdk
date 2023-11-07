package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/zscboy/titan-game-sdk/storage/memfile"
)

func TestCalculateCarCID(t *testing.T) {
	f, err := os.Open("./example/main.go")
	if err != nil {
		t.Fatal(err)
	}

	cid, err := CalculateCid(f)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("cid ", cid.String())
}

func TestCreateCarWithFile(t *testing.T) {
	// }
	input := "./example/example.exe"
	output := "./example/example.car"

	root, err := createCar(input, output)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("root ", root.String())

}

func TestCreateCarStream(t *testing.T) {
	f, err := os.Open("./example/example.exe")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	// carFile, err := os.Create("./example/example.car")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// defer carFile.Close()

	mFile := memfile.New([]byte{})
	root, err := createCarStream(f, mFile)
	if err != nil {
		t.Fatal(err)
	}

	mFile.Seek(0, 0)

	// stat, err := carFile.Stat()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// buf := mFile.Bytes()
	t.Log("car file size:", len(mFile.Bytes()))

	testFile, err := os.Create("./example/test")
	if err != nil {
		t.Fatal(err)
	}
	defer testFile.Close()

	io.Copy(testFile, mFile)

	stat, err := testFile.Stat()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("test file size:", stat.Size())
	t.Log("root ", root.String())
}

func TestUpload(t *testing.T) {
	const apiKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJ1c2VyIl0sIklEIjoiMTA1MjQ0MTYwN0BxcS5jb20iLCJOb2RlSUQiOiIiLCJFeHRlbmQiOiIifQ.Yjoxg9JA7SuikMFL0hHMtOANH1CD2v3JKbpkhSC88XQ"
	const locatorURL = "https://120.79.221.36:5000/rpc/v0"
	storage, close, err := NewStorage(locatorURL, apiKey)
	if err != nil {
		t.Fatal("NewStorage error ", err)
	}
	defer close()

	progress := func(doneSize int64, totalSize int64) {
		t.Logf("upload %d of %d", doneSize, totalSize)
	}

	filePath := "./"
	visitFile := func(fp string, fi os.DirEntry, err error) error {
		// Check for and handle errors
		if err != nil {
			fmt.Println(err) // Can be used to handle errors (e.g., permission denied)
			return nil
		}
		if fi.IsDir() {
			return nil
		} else {
			// This is a file, you can perform file-specific operations here
			if strings.HasSuffix(fp, ".go") {
				path, err := filepath.Abs(fp)
				if err != nil {
					t.Fatal(err)
				}
				_, err = storage.UploadFilesWithPath(context.Background(), path, progress)
				if err != nil {
					t.Log("upload file failed ", err.Error())
					return nil
				}

				t.Logf("totalSize %s success", fp)
			}

		}
		return nil
	}

	err = filepath.WalkDir(filePath, visitFile)
	if err != nil {
		t.Fatal("WalkDir ", err)
	}
}

func TestUploadStream(t *testing.T) {
	const apiKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJ1c2VyIl0sIklEIjoiMTA1MjQ0MTYwN0BxcS5jb20iLCJOb2RlSUQiOiIiLCJFeHRlbmQiOiIifQ.Yjoxg9JA7SuikMFL0hHMtOANH1CD2v3JKbpkhSC88XQ"
	const locatorURL = "https://120.79.221.36:5000/rpc/v0"
	storage, close, err := NewStorage(locatorURL, apiKey)
	if err != nil {
		t.Fatal("NewStorage error ", err)
	}
	defer close()

	progress := func(doneSize int64, totalSize int64) {
		t.Logf("upload %d of %d", doneSize, totalSize)
	}

	filePath := "./storage_test.go"
	f, err := os.Open(filePath)
	if err != nil {
		t.Fatal(err)
	}

	cid, err := storage.UploadStream(context.Background(), f, f.Name(), progress)
	if err != nil {
		t.Fatal("upload file failed ", err.Error())
	}

	t.Logf("totalSize %s success, cid %s", filePath, cid.String())

}

func TestGetFile(t *testing.T) {
	const apiKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJ1c2VyIl0sIklEIjoiMTA1MjQ0MTYwN0BxcS5jb20iLCJOb2RlSUQiOiIiLCJFeHRlbmQiOiIifQ.Yjoxg9JA7SuikMFL0hHMtOANH1CD2v3JKbpkhSC88XQ"
	const locatorURL = "https://120.79.221.36:5000/rpc/v0"
	s, close, err := NewStorage(locatorURL, apiKey)
	if err != nil {
		t.Fatal("NewStorage error ", err)
	}
	defer close()

	storageObject := s.(*storage)
	t.Log("candidate node ", storageObject.candidateID)

	progress := func(doneSize int64, totalSize int64) {
		t.Logf("upload %d of %d", doneSize, totalSize)
	}

	filePath := "./storage_test.go"
	f, err := os.Open(filePath)
	if err != nil {
		t.Fatal(err)
	}

	cid, err := s.UploadStream(context.Background(), f, f.Name(), progress)
	if err != nil {
		t.Fatal("upload file failed ", err.Error())
	}

	url, err := s.GetURL(context.Background(), cid.String())
	if err != nil {
		t.Fatal("get url ", err)
	}

	t.Log("url:", url)

	reader, err := s.GetFileWithCid(context.Background(), cid.String())
	if err != nil {
		t.Fatal("get url ", err)
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal("get url ", err)
	}

	t.Logf("get file %s %d", cid.String(), len(data))
}

func TestUploadFileWithURL(t *testing.T) {
	const apiKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJ1c2VyIl0sIklEIjoiMTA1MjQ0MTYwN0BxcS5jb20iLCJOb2RlSUQiOiIiLCJFeHRlbmQiOiIifQ.Yjoxg9JA7SuikMFL0hHMtOANH1CD2v3JKbpkhSC88XQ"
	const locatorURL = "https://120.79.221.36:5000/rpc/v0"
	s, close, err := NewStorage(locatorURL, apiKey)
	if err != nil {
		t.Fatal("NewStorage error ", err)
	}
	defer close()

	url := "https://files.oaiusercontent.com/file-HQiDjktehYWarlxwnUNF7djs?se=2023-11-07T09%3A38%3A06Z&sp=r&sv=2021-08-06&sr=b&rscc=max-age%3D31536000%2C%20immutable&rscd=attachment%3B%20filename%3D563c233a-76a1-4c85-9d07-a641e1e5937a.webp&sig=rJBb9kHyciaBUQn%2BQmnutv%2B0W3EqOwU5uWBrbHUHtdc%3D"
	name, err := getFileNameFromURL(url)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("name:", name)
	cid, newURL, err := s.UploadFileWithURL(context.Background(), url, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("cid %s, newURL %s", cid, newURL)

}
