package xunlei

import (
	"context"
	"io"
	"net/url"
	"strings"

	"github.com/kkkunny/xunlei/internal/api"
)

// GetMagentFromFile 从文件中获取磁链
func (cli *Client) GetMagentFromFile(ctx context.Context, filename string, file io.Reader) (*url.URL, error) {
	var resp *api.BitTorrentInfoResponse
	err := cli.requestWithCheckAuth(ctx, func() (err error) {
		resp, err = api.BitTorrentInfo(ctx, cli.addr, &api.BitTorrentInfoRequest{
			PanAuth:  cli.panAuth,
			FileName: filename,
			File:     file,
		})
		if err != nil && strings.Contains(err.Error(), "402 Payment Required") {
			return errPanAuthExpired
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return url.Parse(resp.URL)
}
