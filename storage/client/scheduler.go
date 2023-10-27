package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type AssetProperty struct {
	AssetCID  string
	AssetName string
	AssetSize int64
	AssetType string
	NodeID    string
}

type CreateAssetRsp struct {
	UploadURL     string
	Token         string
	AlreadyExists bool
}

type Scheduler interface {
	CreateUserAsset(ctx context.Context, ap *AssetProperty) (*CreateAssetRsp, error)
	DeleteUserAsset(ctx context.Context, assetCID string) error
	ShareUserAssets(ctx context.Context, assetCID []string) (map[string]string, error)
}

var _ Scheduler = (*scheduler)(nil)

func NewScheduler(url string, header http.Header, opts ...Option) Scheduler {
	options := []Option{URLOption(url), HeaderOption(header)}
	options = append(options, opts...)

	client := NewClient(options...)

	return &scheduler{client: client}
}

type scheduler struct {
	client *Client
}

func (s *scheduler) CreateUserAsset(ctx context.Context, ap *AssetProperty) (*CreateAssetRsp, error) {
	serializedParams := params{
		ap,
	}

	req := request{
		Jsonrpc: "2.0",
		Method:  "titan.CreateUserAsset",
		Params:  serializedParams,
		ID:      1,
	}

	rsp, err := s.client.request(ctx, req)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(rsp.Result)
	if err != nil {
		return nil, err
	}

	if rsp.Error != nil {
		return nil, fmt.Errorf("%s code %d ", rsp.Error.Message, rsp.Error.Code)
	}

	createAssetRsp := &CreateAssetRsp{}
	err = json.Unmarshal(b, &createAssetRsp)
	if err != nil {
		return nil, err
	}

	return createAssetRsp, nil
}

func (s *scheduler) DeleteUserAsset(ctx context.Context, assetCID string) error {
	serializedParams := params{
		assetCID,
	}

	req := request{
		Jsonrpc: "2.0",
		Method:  "titan.DeleteUserAsset",
		Params:  serializedParams,
		ID:      1,
	}

	rsp, err := s.client.request(ctx, req)
	if err != nil {
		return err
	}

	if rsp.Error != nil {
		return fmt.Errorf("%s code %d ", rsp.Error.Message, rsp.Error.Code)
	}

	return nil

}

func (s *scheduler) ShareUserAssets(ctx context.Context, assetCID []string) (map[string]string, error) {
	serializedParams := params{
		assetCID,
	}

	req := request{
		Jsonrpc: "2.0",
		Method:  "titan.ShareUserAssets",
		Params:  serializedParams,
		ID:      1,
	}

	rsp, err := s.client.request(ctx, req)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(rsp.Result)
	if err != nil {
		return nil, err
	}

	if rsp.Error != nil {
		return nil, fmt.Errorf("%s code %d ", rsp.Error.Message, rsp.Error.Code)
	}

	ret := make(map[string]string)
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
