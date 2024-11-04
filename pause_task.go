package xunlei

import (
	"context"

	"github.com/kkkunny/xunlei/dto"
)

// PauseTask 暂停任务
func (cli *Client) PauseTask(ctx context.Context, taskID string) error {
	return cli.ModifyTaskPhase(ctx, taskID, dto.TaskPhaseTypePaused)
}
