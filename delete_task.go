package xunlei

import (
	"context"
	"strings"

	"github.com/kkkunny/xunlei/dto"
	"github.com/kkkunny/xunlei/internal/api"
)

// DeleteTask 删除任务
func (cli *Client) DeleteTask(ctx context.Context, taskID string, withLocalFile bool) error {
	if withLocalFile {
		return cli.ModifyTaskPhase(ctx, taskID, dto.TaskPhaseTypeDelete)
	}

	return cli.requestWithCheckAuth(ctx, func() error {
		err := api.DeleteTask(ctx, cli.addr, &api.DeleteTaskRequest{
			Space:   cli.getSpace(),
			PanAuth: cli.panAuth,
			TaskIDs: []string{taskID},
		})
		if err != nil && strings.Contains(err.Error(), "402 Payment Required") {
			return errPanAuthExpired
		}
		return err
	})
}
