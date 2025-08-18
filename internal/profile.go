package im

import "context"

func (c *Client) PortraitSet(ctx context.Context, user string, items []struct {
	Tag   string
	Value interface{}
}) ([]byte, int, error) {
	req := PortraitSetReq{FromAccount: user}
	for _, it := range items {
		req.ProfileItem = append(req.ProfileItem, struct {
			Tag   string      `json:"Tag"`
			Value interface{} `json:"Value"`
		}{Tag: it.Tag, Value: it.Value})
	}
	return c.postJSON(ctx, PortraitSetAPIPath, req)
}

func (c *Client) PortraitGet(ctx context.Context, users []string, tags []string) ([]byte, int, error) {
	return c.postJSON(ctx, PortraitGetAPIPath, PortraitGetReq{ToAccount: users, TagList: tags})
}
