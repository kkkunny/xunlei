package xunlei

import (
	"context"

	"github.com/samber/lo"

	"github.com/kkkunny/xunlei/dto"
	"github.com/kkkunny/xunlei/internal/api"
	"github.com/kkkunny/xunlei/internal/api/conv"
)

// ListTasks 获取任务列表
func (cli *Client) ListTasks(ctx context.Context, allowPhases ...dto.TaskPhase) ([]*dto.TaskInfo, error) {
	resp, err := api.ListTasks(ctx, cli.addr, &api.ListTasksRequest{
		Space: cli.getSpace(),
		Limit: 1000,
		Filter: &api.ListTasksFilter{
			AllowPhases: lo.Map(allowPhases, func(phase dto.TaskPhase, _ int) string {
				return string(phase)
			}),
		},
		PanAuth: cli.panAuth,
	})
	if err != nil {
		return nil, err
	}
	tasks := make([]*dto.TaskInfo, len(resp.Tasks))
	for i, t := range resp.Tasks {
		tasks[i], err = conv.ConvTaskInfoToDTO(t)
		if err != nil {
			return nil, err
		}
	}
	return tasks, nil
}
