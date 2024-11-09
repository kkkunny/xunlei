package xunlei

import (
	"context"
	"errors"
	"fmt"

	"github.com/kkkunny/xunlei/internal/api"
)

var errPanAuthExpired = errors.New("pan-auth expired")

func (cli *Client) requestWithCheckAuth(ctx context.Context, f func() error) error {
	err := f()
	if err == nil || !errors.Is(err, errPanAuthExpired) {
		return err
	}

	panAuth, err := api.GetPanAuth(ctx, cli.addr)
	if err != nil {
		return err
	}
	cli.panAuth = panAuth

	err = f()
	if err == nil || !errors.Is(err, errPanAuthExpired) {
		return err
	}
	return fmt.Errorf("can not get valid pan-auth")
}
