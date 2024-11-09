package xunlei

import (
	"context"
	"strings"

	"github.com/kkkunny/xunlei/dto"
	"github.com/kkkunny/xunlei/internal/api"
	"github.com/kkkunny/xunlei/internal/api/conv"
)

// ListResource 列出远程资源
func (cli *Client) ListResource(ctx context.Context, url string) ([]dto.Resource, error) {
	var resp *api.ListResourceResponse
	err := cli.requestWithCheckAuth(ctx, func() (err error) {
		resp, err = api.ListResource(ctx, cli.addr, &api.ListResourceRequest{
			PanAuth:  cli.panAuth,
			PageSize: 1000,
			URL:      url,
		})
		if err != nil && strings.Contains(err.Error(), "402 Payment Required") {
			return errPanAuthExpired
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	resources := make([]dto.Resource, len(resp.List.Resources))
	for i, r := range resp.List.Resources {
		resources[i], err = conv.ConvResourceInfoToDTO(r)
		if err != nil {
			return nil, err
		}
	}
	return resources, nil
}
