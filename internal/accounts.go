package im

import "context"

func (c *Client) AccountImport(ctx context.Context, id, nick, face string) ([]byte, int, error) {
	return c.postJSON(ctx, AccountImportAPIPath, AccountImportReq{UserID: id, Nick: nick, FaceUrl: face})
}
func (c *Client) AccountCheck(ctx context.Context, ids ...string) ([]byte, int, error) {
	items := make([]struct {
		UserID string `json:"UserID"`
	}, 0, len(ids))
	for _, id := range ids {
		items = append(items, struct {
			UserID string `json:"UserID"`
		}{UserID: id})
	}
	return c.postJSON(ctx, AccountCheckAPIPath, AccountCheckReq{CheckItem: items})
}
func (c *Client) Kick(ctx context.Context, id string) ([]byte, int, error) {
	return c.postJSON(ctx, KickAPIPath, KickReq{UserID: id})
}
func (c *Client) QueryOnlineStatus(ctx context.Context, ids []string, detail, instid, custom bool) ([]byte, int, error) {
	req := QueryOnlineStatusReq{ToAccount: ids}
	if detail {
		req.IsNeedDetail = 1
	}
	if instid {
		req.IsReturnInstid = 1
	}
	if custom {
		req.IsReturnCustomStatus = 1
	}
	return c.postJSON(ctx, QueryOnlineStatusAPIPath, req)
}

func (c *Client) MultiAccountImport(ctx context.Context, accounts []AccountImportReq) ([]byte, int, error) {
	return c.postJSON(ctx, MultiAccountImportAPIPath, MultiAccountImportReq{Accounts: accounts})
}

func (c *Client) AccountDelete(ctx context.Context, ids ...string) ([]byte, int, error) {
	items := make([]struct {
		UserID string `json:"UserID"`
	}, 0, len(ids))
	for _, id := range ids {
		if id == "" {
			continue
		}
		items = append(items, struct {
			UserID string `json:"UserID"`
		}{UserID: id})
	}
	return c.postJSON(ctx, AccountDeleteAPIPath, AccountDeleteReq{DeleteItem: items})
}
