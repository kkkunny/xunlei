package xunlei

import (
	"context"
	"io"
	"net/url"

	"github.com/kkkunny/xunlei/internal/api"
)

// GetMagentFromFile 从文件中获取磁链
func (cli *Client) GetMagentFromFile(ctx context.Context, filename string, file io.Reader) (*url.URL, error) {
	resp, err := api.BitTorrentInfo(ctx, cli.addr, &api.BitTorrentInfoRequest{
		PanAuth:  cli.panAuth,
		FileName: filename,
		File:     file,
	})
	if err != nil {
		return nil, err
	}
	return url.Parse(resp.URL)
}
