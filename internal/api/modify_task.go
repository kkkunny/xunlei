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

type PatchTaskRequest struct {
	Space     string           `json:"space"`
	Type      string           `json:"type"`
	ID        string           `json:"id"`
	PanAuth   string           `json:"pan_auth"`
	SetParams *PatchTaskParams `json:"set_params"`
}

type PatchTaskParams struct {
	Phase string `json:"phase"`
}

type PatchTaskResponse struct {
	HttpStatus int64 `json:"HttpStatus"`
}

// PatchTask 修改任务
func PatchTask(ctx context.Context, addr string, req *PatchTaskRequest) (*PatchTaskResponse, error) {
	path, err := url.JoinPath(addr, "webman", "3rdparty", "pan-xunlei-com", "index.cgi", "method", "patch", "drive", "v1", "task")
	if err != nil {
		return nil, err
	}

	var setParams string
	if req.SetParams != nil {
		setParamData, err := json.Marshal(req.SetParams)
		if err != nil {
			return nil, err
		}
		setParams = string(setParamData)
	}

	var bizResp PatchTaskResponse
	resp, err := request.C().SetTimeout(time.Second*5).R().SetContext(ctx).
		SetQueryParam("pan_auth", req.PanAuth).
		SetBodyJsonMarshal(map[string]any{
			"space": req.Space,
			"type":  req.Type,
			"id":    req.ID,
			"set_params": map[string]string{
				"spec": setParams,
			},
		}).
		SetSuccessResult(&bizResp).
		Post(path)
	if err != nil {
		return nil, err
	} else if resp.GetStatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unknown http error: StatusCode=%d, Status=%s", resp.GetStatusCode(), resp.GetStatus())
	}
	return &bizResp, nil
}
