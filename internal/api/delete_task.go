package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	request "github.com/imroc/req/v3"
)

type DeleteTaskRequest struct {
	Space   string   `json:"space"`
	PanAuth string   `json:"pan_auth"`
	TaskIDs []string `json:"task_ids"`
}

// DeleteTask 删除任务（不会同时删除本地文件，如要删除文件，请使用ModifyTask）
func DeleteTask(ctx context.Context, addr string, req *DeleteTaskRequest) error {
	path, err := url.JoinPath(addr, "webman", "3rdparty", "pan-xunlei-com", "index.cgi", "method", "delete", "drive", "v1", "tasks")
	if err != nil {
		return err
	}

	resp, err := request.C().SetTimeout(time.Second * 5).R().SetContext(ctx).
		SetQueryParams(map[string]string{
			"space":    req.Space,
			"task_ids": strings.Join(req.TaskIDs, ","),
			"pan_auth": req.PanAuth,
		}).
		Post(path)
	if err != nil {
		return err
	} else if resp.GetStatusCode() != http.StatusOK {
		return fmt.Errorf("unknown http error: StatusCode=%d, Status=%s", resp.GetStatusCode(), resp.GetStatus())
	}
	return nil
}
