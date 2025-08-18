package im

import (
	"context"
	"encoding/json"
	"fmt"
)

func (c *Client) SendGroupText(ctx context.Context, groupId, text, from string) (*SendGroupMsgResp, []byte, error) {
	rnd, _ := c.randUint32()
	body := SendGroupMsgReq{GroupId: groupId, Random: rnd, MsgBody: []MsgElem{{MsgType: "TIMTextElem", MsgContent: map[string]any{"Text": text}}}}
	if from != "" {
		body.FromAccount = from
	}
	b, code, err := c.postJSON(ctx, SendGroupMsgAPIPath, body)
	if err != nil {
		return nil, b, fmt.Errorf("HTTP %d: %w", code, err)
	}
	var res SendGroupMsgResp
	_ = json.Unmarshal(b, &res)
	return &res, b, nil
}
func (c *Client) GetGroupHistory(ctx context.Context, groupId string, n int, beforeSeq uint64, withRecalled bool) ([]byte, int, error) {
	req := GroupMsgGetSimpleReq{GroupId: groupId, ReqMsgNumber: n}
	if beforeSeq > 0 {
		req.ReqMsgSeq = beforeSeq
	}
	if withRecalled {
		req.WithRecalledMsg = 1
	}
	return c.postJSON(ctx, GroupMsgGetSimpleAPIPath, req)
}
func (c *Client) RecallGroupMsgs(ctx context.Context, groupId string, reason string, seqs ...uint64) ([]byte, int, error) {
	list := make([]struct {
		MsgSeq uint64 `json:"MsgSeq"`
	}, 0, len(seqs))
	for _, s := range seqs {
		list = append(list, struct {
			MsgSeq uint64 `json:"MsgSeq"`
		}{s})
	}
	return c.postJSON(ctx, GroupMsgRecallAPIPath, GroupMsgRecallReq{GroupId: groupId, MsgSeqList: list, Reason: reason})
}
func (c *Client) SendC2CText(ctx context.Context, to, text, from string, sync bool) (*SendC2CMsgResp, []byte, error) {
	rnd, _ := c.randUint32()
	req := SendC2CMsgReq{ToAccount: to, MsgRandom: rnd, MsgBody: []MsgElem{{MsgType: "TIMTextElem", MsgContent: map[string]any{"Text": text}}}}
	if from != "" {
		req.FromAccount = from
	}
	if !sync {
		req.SyncOtherMachine = 2
	} else {
		req.SyncOtherMachine = 1
	}
	b, code, err := c.postJSON(ctx, SendC2CMsgAPIPath, req)
	if err != nil {
		return nil, b, fmt.Errorf("HTTP %d: %w", code, err)
	}
	var res SendC2CMsgResp
	_ = json.Unmarshal(b, &res)
	return &res, b, nil
}

// BatchSendC2CText: 여러 대상에게 1:1 텍스트 전송
func (c *Client) BatchSendC2CText(ctx context.Context, to []string, text, from string, sync bool) ([]byte, int, error) {
	rnd, _ := c.randUint32()
	req := BatchSendC2CMsgReq{
		ToAccount: to,
		MsgRandom: rnd,
		MsgBody:   []MsgElem{{MsgType: "TIMTextElem", MsgContent: map[string]any{"Text": text}}},
	}
	if from != "" {
		req.FromAccount = from
	}
	if !sync {
		req.SyncOtherMachine = 2
	} else {
		req.SyncOtherMachine = 1
	}
	return c.postJSON(ctx, BatchSendC2CMsgAPIPath, req)
}

// AdminSetMsgRead: 대화 읽음 처리
func (c *Client) AdminSetMsgRead(ctx context.Context, report, peer string, readTime uint64) ([]byte, int, error) {
	req := AdminSetMsgReadReq{ReportAccount: report, PeerAccount: peer}
	if readTime > 0 {
		req.MsgReadTime = readTime
	}
	return c.postJSON(ctx, AdminSetMsgReadAPIPath, req)
}

// AdminMsgWithdraw: MsgKey 기반 C2C 메시지 회수
func (c *Client) AdminMsgWithdraw(ctx context.Context, from, to, msgKey string) ([]byte, int, error) {
	return c.postJSON(ctx, AdminMsgWithdrawAPIPath, AdminMsgWithdrawReq{FromAccount: from, ToAccount: to, MsgKey: msgKey})
}

// GetC2CUnreadMsgNum: 미읽음 개수 조회
func (c *Client) GetC2CUnreadMsgNum(ctx context.Context, to string, peers []string) ([]byte, int, error) {
	return c.postJSON(ctx, GetC2CUnreadMsgNumAPIPath, GetC2CUnreadMsgNumReq{ToAccount: to, PeerAccount: peers})
}
