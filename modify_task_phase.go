package xunlei

import (
	"context"
	"strings"

	"github.com/kkkunny/xunlei/dto"
	"github.com/kkkunny/xunlei/internal/api"
)

// Deprecated: 请使用 ContinueTask、PauseTask、DeleteTask 代替
// ModifyTaskPhase 修改任务状态
func (cli *Client) ModifyTaskPhase(ctx context.Context, taskID string, phase dto.TaskPhase) error {
	return cli.requestWithCheckAuth(ctx, func() error {
		_, err := api.PatchTask(ctx, cli.addr, &api.PatchTaskRequest{
			Space:   cli.getSpace(),
			PanAuth: cli.panAuth,
			ID:      taskID,
			Param: api.PatchTaskParam{
				Phase: phase.Spec(),
			},
		})
		if err != nil && strings.Contains(err.Error(), "402 Payment Required") {
			return errPanAuthExpired
		}
		return err
	})
}
