package xunlei

import (
	"context"

	"github.com/kkkunny/xunlei/dto"
)

// ContinueTask 继续任务
func (cli *Client) ContinueTask(ctx context.Context, taskID string) error {
	return cli.ModifyTaskPhase(ctx, taskID, dto.TaskPhaseTypeRunning)
}
