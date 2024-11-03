package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	request "github.com/imroc/req/v3"
)

type ListTasksRequest struct {
	Space   string            `json:"space"`
	Limit   int64             `json:"limit"`
	Filters *ListTasksFilters `json:"filters"`
	PanAuth string            `json:"pan_auth"`
}

type ListTasksFilters struct {
	AllowPhases []string `json:"phase"`
	AllowTypes  []string `json:"type"`
}

type ListTasksResponse struct {
	HttpStatus int64      `json:"HttpStatus"`
	Tasks      []TaskInfo `json:"tasks"`
	ExpiresIn  int64      `json:"expires_in"`
}

type TaskInfo struct {
	Kind        string            `json:"kind"`
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Type        string            `json:"type"`
	UserID      string            `json:"user_id"`
	Params      map[string]string `json:"params"`
	FileID      string            `json:"file_id,omitempty"`
	FileName    string            `json:"file_name"`
	FileSize    string            `json:"file_size"`
	Message     string            `json:"message,omitempty"`
	CreatedTime string            `json:"created_time"`
	UpdatedTime string            `json:"updated_time"`
	IconLink    string            `json:"icon_link"`
	Phase       string            `json:"phase"`
	Progress    int64             `json:"progress,omitempty"`
	Space       string            `json:"space"`
}

// ListTasks 查看所有任务
func ListTasks(ctx context.Context, addr string, req *ListTasksRequest) (*ListTasksResponse, error) {
	path, err := url.JoinPath(addr, "webman", "3rdparty", "pan-xunlei-com", "index.cgi", "drive", "v1", "tasks")
	if err != nil {
		return nil, err
	}

	var filters string
	if req.Filters != nil {
		filterMap := make(map[string]any)
		if len(req.Filters.AllowPhases) > 0 {
			filterMap["phase"] = map[string]string{
				"in": strings.Join(req.Filters.AllowPhases, ","),
			}
		}
		if len(req.Filters.AllowTypes) > 0 {
			filterMap["type"] = map[string]string{
				"in": strings.Join(req.Filters.AllowTypes, ","),
			}
		}
		filterData, err := json.Marshal(filterMap)
		if err != nil {
			return nil, err
		}
		filters = string(filterData)
	}

	var bizResp ListTasksResponse
	resp, err := request.C().SetTimeout(time.Second * 5).R().SetContext(ctx).
		SetQueryParams(map[string]string{
			"space":    req.Space,
			"limit":    strconv.FormatInt(req.Limit, 10),
			"filters":  filters,
			"pan_auth": req.PanAuth,
		}).
		SetSuccessResult(&bizResp).
		Get(path)
	if err != nil {
		return nil, err
	} else if resp.GetStatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unknown http error: StatusCode=%d, Status=%s", resp.GetStatusCode(), resp.GetStatus())
	}
	return &bizResp, nil
}
