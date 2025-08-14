package im

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) CreateGroup(ctx context.Context, groupId, name, owner, groupType string) (*ResponseData, []byte, error) {
	if name == "" {
		name = groupId
	}
	body := ChatGroupBody{OwnerAccount: owner, Type: groupType, GroupId: groupId, Name: name}
	b, code, err := c.postJSON(ctx, CreateGroupAPIPath, body)
	if err != nil {
		return nil, b, fmt.Errorf("HTTP %d: %w", code, err)
	}
	var res ResponseData
	_ = json.Unmarshal(b, &res)
	return &res, b, nil
}
func (c *Client) DestroyGroup(ctx context.Context, groupId string) (*ResponseData, []byte, error) {
	b, code, err := c.postJSON(ctx, DestroyGroupAPIPath, DestroyGroupReq{GroupId: groupId})
	if err != nil {
		return nil, b, fmt.Errorf("HTTP %d: %w", code, err)
	}
	var res ResponseData
	_ = json.Unmarshal(b, &res)
	return &res, b, nil
}
func (c *Client) GetAllGroups(ctx context.Context, limit int, next uint64) (*ResponseData, []byte, error) {
	b, code, err := c.postJSON(ctx, GetAllGroupsAPIPath, GetGroupsRequest{Limit: limit, Next: next})
	if err != nil {
		return nil, b, fmt.Errorf("HTTP %d: %w", code, err)
	}
	var res ResponseData
	_ = json.Unmarshal(b, &res)
	return &res, b, nil
}
func (c *Client) SendGroupSystemNotification(ctx context.Context, groupId, content string, toMembers []string) (*ResponseData, []byte, error) {
	body := GroupSystemNotificationReq{GroupId: groupId, Content: content}
	if len(toMembers) > 0 {
		body.ToMembersAccount = toMembers
	}
	b, code, err := c.postJSON(ctx, SendGroupSystemNotifyPath, body)
	if err != nil {
		return nil, b, fmt.Errorf("HTTP %d: %w", code, err)
	}
	var res ResponseData
	_ = json.Unmarshal(b, &res)
	return &res, b, nil
}
func (c *Client) GetGroupInfo(ctx context.Context, ids ...string) ([]byte, int, error) {
	return c.postJSON(ctx, GetGroupInfoAPIPath, GetGroupInfoReq{GroupIdList: ids})
}
func (c *Client) GetGroupMembers(ctx context.Context, groupId string, limit, offset int) ([]byte, int, error) {
	return c.postJSON(ctx, GetGroupMemberInfoAPIPath, GetGroupMemberInfoReq{GroupId: groupId, Limit: limit, Offset: offset})
}
func (c *Client) AddGroupMembers(ctx context.Context, groupId string, members []string, silent bool) ([]byte, int, error) {
	list := make([]addMember, 0, len(members))
	for _, m := range members {
		if v := strings.TrimSpace(m); v != "" {
			list = append(list, addMember{Member_Account: v})
		}
	}
	req := AddGroupMemberReq{GroupId: groupId, MemberList: list}
	if silent {
		req.Silence = 1
	}
	return c.postJSON(ctx, AddGroupMemberAPIPath, req)
}
func (c *Client) DeleteGroupMembers(ctx context.Context, groupId string, members []string, reason string, silent bool) ([]byte, int, error) {
	req := DeleteGroupMemberReq{GroupId: groupId, MemberToDel_Account: members}
	if reason != "" {
		req.Reason = reason
	}
	if silent {
		req.Silence = 1
	}
	return c.postJSON(ctx, DeleteGroupMemberAPIPath, req)
}
func (c *Client) MuteMembers(ctx context.Context, groupId string, members []string, seconds uint32) ([]byte, int, error) {
	return c.postJSON(ctx, ForbidSendMsgAPIPath, ForbidSendMsgReq{GroupId: groupId, Members_Account: members, MuteTime: seconds})
}
