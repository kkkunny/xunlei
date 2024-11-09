package xunlei

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"

	"github.com/kkkunny/xunlei/dto"
	"github.com/kkkunny/xunlei/internal/api"
	"github.com/kkkunny/xunlei/internal/api/conv"
)

// CreateTask 创建任务
func (cli *Client) CreateTask(ctx context.Context, name string, url string, subFileFilter ...func(file *dto.FileResource) bool) (*dto.TaskInfo, error) {
	resources, err := cli.ListResource(ctx, url)
	if err != nil {
		return nil, err
	} else if len(resources) == 0 {
		return nil, fmt.Errorf("url not found resources")
	}
	resource := resources[0]

	if name == "" {
		name = resource.GetName()
	}

	var subFileIndex string
	if len(subFileFilter) == 0 {
		subFileIndex = fmt.Sprintf("0-%d", resource.GetFileCount()-1)
	} else {
		subFileFilterFn := subFileFilter[len(subFileFilter)-1]
		files := lo.Filter(resource.GetFiles(), func(file *dto.FileResource, _ int) bool {
			return subFileFilterFn(file)
		})
		if len(files) == 0 {
			return nil, fmt.Errorf("no file will download")
		} else {
			fileIndexes := lo.Map(files, func(file *dto.FileResource, _ int) int64 {
				return file.FileIndex
			})
			slices.Sort(fileIndexes)
			if fileIndexes[len(fileIndexes)-1]-fileIndexes[0] == int64(len(fileIndexes))-1 {
				subFileIndex = fmt.Sprintf("%d-%d", fileIndexes[0], fileIndexes[len(fileIndexes)-1])
			} else {
				subFileIndex = strings.Join(lo.Map(fileIndexes, func(index int64, _ int) string { return strconv.FormatInt(index, 10) }), ",")
			}
		}
	}

	var resp *api.CreateTaskResponse
	err = cli.requestWithCheckAuth(ctx, func() (err error) {
		resp, err = api.CreateTask(ctx, cli.addr, &api.CreateTaskRequest{
			PanAuth:  cli.panAuth,
			Type:     string(dto.TaskTypeUserDownloadURL),
			Space:    cli.getSpace(),
			Name:     name,
			FileName: resource.GetName(),
			FileSize: resource.GetFileSize(),
			Param: api.CreateTaskParam{
				Target:       cli.getSpace(),
				URL:          url,
				SubFileIndex: subFileIndex,
			},
		})
		if err != nil && strings.Contains(err.Error(), "402 Payment Required") {
			return errPanAuthExpired
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return conv.ConvTaskInfoToDTO(&resp.Task)
}
