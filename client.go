package xunlei

import "fmt"

type Client struct {
	addr     string
	deviceID string
	panAuth  string
}

func NewClient(addr string, did string) *Client {
	return &Client{
		addr:     addr,
		deviceID: did,
	}
}

func (cli *Client) getSpace() string {
	return fmt.Sprintf("device_id#%s", cli.deviceID)
}
