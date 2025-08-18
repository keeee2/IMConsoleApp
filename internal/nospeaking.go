package im

import "context"

func (c *Client) SetNoSpeaking(ctx context.Context, user string, c2c, group uint32) ([]byte, int, error) {
	return c.postJSON(ctx, SetNoSpeakingAPIPath, SetNoSpeakingReq{Set_Account: user, C2CmsgNospeakingTime: c2c, GroupmsgNospeakingTime: group})
}

func (c *Client) GetNoSpeaking(ctx context.Context, user string) ([]byte, int, error) {
	return c.postJSON(ctx, GetNoSpeakingAPIPath, GetNoSpeakingReq{Get_Account: user})
}
