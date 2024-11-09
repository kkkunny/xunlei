package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"time"

	request "github.com/imroc/req/v3"
)

// GetPanAuth 获取pan-auth
func GetPanAuth(ctx context.Context, addr string) (string, error) {
	path, err := url.JoinPath(addr, "webman", "3rdparty", "pan-xunlei-com", "index.cgi/")
	if err != nil {
		return "", err
	}

	resp, err := request.C().SetTimeout(time.Second * 5).R().SetContext(ctx).Get(path)
	if err != nil {
		return "", err
	} else if resp.GetStatusCode() != http.StatusOK {
		return "", fmt.Errorf("unknown http error: StatusCode=%d, Status=%s", resp.GetStatusCode(), resp.GetStatus())
	}
	panAuthRes := regexp.MustCompile(`function\s*uiauth\s*\(\s*value\s*\)\s*\{\s*return\s*"(.+?)"\s*}`).FindStringSubmatch(resp.String())
	if len(panAuthRes) != 2 {
		return "", fmt.Errorf("can not found pan-auth")
	}
	return panAuthRes[1], nil
}
