package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	request "github.com/imroc/req/v3"
)

type BitTorrentInfoRequest struct {
	PanAuth  string    `json:"pan_auth"`
	FileName string    `json:"file_name"`
	File     io.Reader `json:"file"`
}

type BitTorrentInfoResponse struct {
	Error string `json:"error"`
	URL   string `json:"url"`
}

// BitTorrentInfo 根据文件解出磁链
func BitTorrentInfo(ctx context.Context, addr string, req *BitTorrentInfoRequest) (*BitTorrentInfoResponse, error) {
	path, err := url.JoinPath(addr, "webman", "3rdparty", "pan-xunlei-com", "index.cgi", "device", "btinfo")
	if err != nil {
		return nil, err
	}

	var bizResp BitTorrentInfoResponse
	resp, err := request.C().SetTimeout(time.Second*5).R().SetContext(ctx).
		SetQueryParam("pan_auth", req.PanAuth).
		SetFileUpload(request.FileUpload{
			ParamName:   "file",
			FileName:    req.FileName,
			ContentType: "application/octet-stream",
			GetFileContent: func() (io.ReadCloser, error) {
				if rc, ok := req.File.(io.ReadCloser); ok {
					return rc, nil
				}
				return io.NopCloser(req.File), nil
			},
		}).
		SetFormData(map[string]string{"pan-auth": req.PanAuth}).
		SetSuccessResult(&bizResp).
		Post(path)
	if err != nil {
		return nil, err
	} else if resp.GetStatusCode() != http.StatusOK {
		return nil, fmt.Errorf("unknown http error: StatusCode=%d, Status=%s", resp.GetStatusCode(), resp.GetStatus())
	}
	return &bizResp, nil
}
