package im

import "context"

func (c *Client) GetAppInfo(ctx context.Context, fields ...string) ([]byte, int, error) {
	return c.postJSON(ctx, GetAppInfoAPIPath, GetAppInfoReq{RequestField: fields})
}

func (c *Client) GetIPList(ctx context.Context) ([]byte, int, error) {
	return c.postJSON(ctx, GetIPListAPIPath, GetIPListReq{})
}
