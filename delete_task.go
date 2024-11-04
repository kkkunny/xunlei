package xunlei

import (
	"context"

	"github.com/kkkunny/xunlei/dto"
	"github.com/kkkunny/xunlei/internal/api"
)

// DeleteTask 删除任务
func (cli *Client) DeleteTask(ctx context.Context, taskID string, withLocalFile ...bool) error {
	isWithLocalFile := false
	if len(withLocalFile) > 0 {
		isWithLocalFile = withLocalFile[len(withLocalFile)]
	}

	if isWithLocalFile {
		return cli.ModifyTaskPhase(ctx, taskID, dto.TaskPhaseTypeDelete)
	}

	err := api.DeleteTask(ctx, cli.addr, &api.DeleteTaskRequest{
		Space:   cli.getSpace(),
		PanAuth: cli.panAuth,
		TaskIDs: []string{taskID},
	})
	return err
}
