package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	request "github.com/imroc/req/v3"
)

type CreateTaskRequest struct {
	PanAuth  string          `json:"pan_auth"`
	FileName string          `json:"file_name,omitempty"`
	FileSize int64           `json:"file_size,string"`
	Name     string          `json:"name"`
	Param    CreateTaskParam `json:"params"`
	Space    string          `json:"space"`
	Type     string          `json:"type"`
}

type CreateTaskParam struct {
	FileID         string `json:"file_id,omitempty"`
	MIMEType       string `json:"mime_type,omitempty"`
	ParentFolderID string `json:"parent_folder_id,omitempty"`
	SubFileIndex   string `json:"sub_file_index,omitempty"`
	Target         string `json:"target"`
	TotalFileCount int64  `json:"total_file_count,omitempty,string"`
	URL            string `json:"url"`
}

type CreateTaskResponse struct {
	HttpStatus       int64    `json:"HttpStatus"`
	Error            string   `json:"error,omitempty"`             // invalid_argument
	ErrorCode        int64    `json:"error_code,omitempty"`        // 3
	ErrorDescription string   `json:"error_description,omitempty"` // 请求参数错误
	Task             TaskInfo `json:"task"`
}

// CreateTask 创建任务
func CreateTask(ctx context.Context, addr string, req *CreateTaskRequest) (*CreateTaskResponse, error) {
	a, _ := json.MarshalIndent(req, "", "  ")
	fmt.Println(string(a))

	path, err := url.JoinPath(addr, "webman", "3rdparty", "pan-xunlei-com", "index.cgi", "drive", "v1", "task")
	if err != nil {
		return nil, err
	}

	var bizResp CreateTaskResponse
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
