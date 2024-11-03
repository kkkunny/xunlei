package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	request "github.com/imroc/req/v3"
)

type ListResourceRequest struct {
	PanAuth  string `json:"pan_auth"`
	PageSize int64  `json:"page_size"`
	URLs     string `json:"urls"`
}

type ListResourceResponse struct {
	ListID string `json:"list_id"`
	List   struct {
		PageSize  int64           `json:"page_size"`
		Resources []*ResourceInfo `json:"resources"`
	} `json:"list"`
}

type ResourceInfo struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	FileSize  int64             `json:"file_size"`
	FileCount int64             `json:"file_count"`
	Meta      map[string]string `json:"meta"`
	ParentID  string            `json:"parent_id,omitempty"`

	// 仅目录有
	IsDir bool `json:"is_dir,omitempty"`
	Dir   struct {
		PageSize  int64           `json:"page_size"`
		Resources []*ResourceInfo `json:"resources"`
	} `json:"dir,omitempty"`

	// 仅文件有
	FileIndex int64 `json:"file_index,omitempty"`
}

// ListResource 列出磁链内资源
func ListResource(ctx context.Context, addr string, req *ListResourceRequest) (*ListResourceResponse, error) {
	path, err := url.JoinPath(addr, "webman", "3rdparty", "pan-xunlei-com", "index.cgi", "drive", "v1", "resource", "list")
	if err != nil {
		return nil, err
	}

	var bizResp ListResourceResponse
	resp, err := request.C().SetTimeout(time.Second*5).R().SetContext(ctx).
		SetQueryParam("pan_auth", req.PanAuth).
		SetBodyJsonMarshal(&req).
		SetSuccessResult(&bizResp).
		Post(path)
	if err != nil {
		return nil, err
	} else if resp.GetStatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unknown http error: StatusCode=%d, Status=%s", resp.GetStatusCode(), resp.GetStatus())
	}
	return &bizResp, nil
}
