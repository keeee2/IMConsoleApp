package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	im "tencent/api/internal"
)

// ========== Handlers: Groups ==========

func handleCreate(ctx context.Context, c *im.Client, r *bufio.Reader, cfg im.Config) {
	fmt.Println("\n[그룹 만들기]")
	gid := prompt(r, "GroupId", "")
	name := gid
	sp := NewSpinner()
	sp.Start("그룹 생성 중")
	res, raw, err := c.CreateGroup(ctx, gid, name, cfg.AdminIdentifier, cfg.DefaultGroupType)
	if err != nil {
		sp.Stopf("그룹 생성 실패: %v", err)
	} else {
		sp.Stopf("그룹 생성 완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
	if res != nil {
		success := fmt.Sprintf("성공: 그룹 생성 (GroupId=%s, Type=%s)", res.GroupId, res.Type)
		printAPISummary(res, success)
	} else {
		printAPISummary(nil, "")
	}
	fmt.Println()
}

func handleDestroy(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[그룹 삭제하기]")
	gid := prompt(r, "삭제할 GroupId", "")
	if gid == "" {
		fmt.Println("GroupId는 필수입니다.")
		return
	}
	sp := NewSpinner()
	sp.Start("그룹 삭제 중")
	res, raw, err := c.DestroyGroup(ctx, gid)
	if err != nil {
		sp.Stopf("그룹 삭제 실패: %v", err)
	} else {
		sp.Stopf("그룹 삭제 완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
	if res != nil {
		printAPISummary(res, "성공: 그룹 삭제 완료")
	} else {
		printAPISummary(nil, "")
	}
	fmt.Println()
}

func handleList(ctx context.Context, c *im.Client, r *bufio.Reader, sess *Session) {
	fmt.Println("\n[그룹 목록 보기]")
	limit := promptInt(r, "Limit(<=10000)", sess.Limit)
	next := promptUint64(r, "Next(다음 페이지 토큰)", sess.Next)
	sp := NewSpinner()
	sp.Start("그룹 목록 불러오는 중")
	res, raw, err := c.GetAllGroups(ctx, limit, next)
	if err != nil {
		sp.Stopf("목록 조회 실패: %v", err)
	} else {
		sp.Stopf("목록 조회 완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
	if res != nil {
		if res.ErrorCode == 0 {
			sess.Next = res.Next
			sess.Limit = limit
		}
		success := fmt.Sprintf("성공: 총 %d개, Next=%d", res.TotalCount, res.Next)
		printAPISummary(res, success)
	} else {
		printAPISummary(nil, "")
	}
	fmt.Println()
}

func handleGetGroupInfo(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[그룹 상세 조회]")
	ids := prompt(r, "GroupIds(쉼표로 구분)", "")
	if strings.TrimSpace(ids) == "" {
		fmt.Println("GroupIds는 필수입니다.")
		return
	}
	idList := strings.Split(ids, ",")
	for i := range idList {
		idList[i] = strings.TrimSpace(idList[i])
	}
	sp := NewSpinner()
	sp.Start("조회 중")
	raw, _, err := c.GetGroupInfo(ctx, idList...)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

func handleGetGroupMembers(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[그룹 멤버 목록]")
	gid := prompt(r, "GroupId", "")
	limit := promptInt(r, "Limit(<=200 권장)", 100)
	offset := promptInt(r, "Offset", 0)
	if gid == "" {
		fmt.Println("GroupId는 필수입니다.")
		return
	}
	sp := NewSpinner()
	sp.Start("조회 중")
	raw, _, err := c.GetGroupMembers(ctx, gid, limit, offset)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

func handleAddMembers(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[그룹 멤버 초대]")
	gid := prompt(r, "GroupId", "")
	line := prompt(r, "Members(쉼표로 구분)", "")
	silent := strings.ToLower(prompt(r, "Silence? (y/N)", "N")) == "y"
	if gid == "" || line == "" {
		fmt.Println("필수값 누락")
		return
	}
	members := strings.Split(line, ",")
	sp := NewSpinner()
	sp.Start("초대 중")
	raw, _, err := c.AddGroupMembers(ctx, gid, members, silent)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

func handleDeleteMembers(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[그룹 멤버 제거]")
	gid := prompt(r, "GroupId", "")
	line := prompt(r, "Members(쉼표로 구분)", "")
	reason := prompt(r, "Reason(선택)", "")
	silent := strings.ToLower(prompt(r, "Silence? (y/N)", "N")) == "y"
	if gid == "" || line == "" {
		fmt.Println("필수값 누락")
		return
	}
	members := strings.Split(line, ",")
	sp := NewSpinner()
	sp.Start("제거 중")
	raw, _, err := c.DeleteGroupMembers(ctx, gid, members, reason, silent)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

func handleMuteMembers(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[그룹 멤버 음소거/해제]")
	gid := prompt(r, "GroupId", "")
	line := prompt(r, "Members(쉼표로 구분)", "")
	sec := uint64(promptInt(r, "MuteTime(sec, 0=해제)", 60))
	if gid == "" || line == "" {
		fmt.Println("필수값 누락")
		return
	}
	members := strings.Split(line, ",")
	sp := NewSpinner()
	sp.Start("반영 중")
	raw, _, err := c.MuteMembers(ctx, gid, members, uint32(sec))
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

// ========== Handlers: Messages ==========

func handleSendSystemMessage(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[시스템 메시지 전송]")
	gid := prompt(r, "GroupId", "")
	content := prompt(r, "Content", "Hello World")
	rec := prompt(r, "ToMembers(쉼표로 구분, 비우면 전체)", "")
	if gid == "" || content == "" {
		fmt.Println("GroupId/Content는 필수입니다.")
		return
	}
	var toMembers []string
	if strings.TrimSpace(rec) != "" {
		parts := strings.Split(rec, ",")
		for _, p := range parts {
			v := strings.TrimSpace(p)
			if v != "" {
				toMembers = append(toMembers, v)
			}
		}
	}
	sp := NewSpinner()
	sp.Start("시스템 메시지 전송 중")
	res, raw, err := c.SendGroupSystemNotification(ctx, gid, content, toMembers)
	if err != nil {
		sp.Stopf("전송 실패: %v", err)
	} else {
		sp.Stopf("전송 완료")
	}
	fmt.Println("\n--- 요청(JSON) ---")
	reqBody, _ := json.MarshalIndent(im.GroupSystemNotificationReq{GroupId: gid, Content: content, ToMembersAccount: toMembers}, "", "  ")
	fmt.Println(string(reqBody))
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
	if res != nil {
		printAPISummary(res, "성공: 시스템 메시지 전송")
	} else {
		printAPISummary(nil, "")
	}
	fmt.Println()
}

func handleSendGroupText(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[그룹 일반 메시지]")
	gid := prompt(r, "GroupId", "")
	text := prompt(r, "Text", "hello world")
	from := prompt(r, "From_Account(빈칸=관리자)", "")
	if gid == "" || text == "" {
		fmt.Println("필수값 누락")
		return
	}
	sp := NewSpinner()
	sp.Start("전송 중")
	res, raw, err := c.SendGroupText(ctx, gid, text, from)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
	if res != nil {
		printAPISummary(&im.ResponseData{ActionStatus: res.ActionStatus, ErrorCode: res.ErrorCode, ErrorInfo: res.ErrorInfo}, "성공: 그룹 텍스트 전송")
	}
}

func handleSendC2CText(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[1:1(C2C) 텍스트]")
	to := prompt(r, "To_Account", "")
	text := prompt(r, "Text", "hello c2c")
	from := prompt(r, "From_Account(빈칸=관리자)", "")
	sync := strings.ToLower(prompt(r, "다른 기기에도 동기화? (Y/n)", "Y")) != "n"
	if to == "" || text == "" {
		fmt.Println("필수값 누락")
		return
	}
	sp := NewSpinner()
	sp.Start("전송 중")
	res, raw, err := c.SendC2CText(ctx, to, text, from, sync)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
	if res != nil {
		printAPISummary(&im.ResponseData{ActionStatus: res.ActionStatus, ErrorCode: res.ErrorCode, ErrorInfo: res.ErrorInfo}, "성공: C2C 텍스트 전송")
	}
}

func handleGroupHistory(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[그룹 히스토리 조회]")
	gid := prompt(r, "GroupId", "")
	n := promptInt(r, "개수(<=20 권장)", 10)
	before := prompt(r, "Before Seq(비우면 최신부터)", "")
	withRec := strings.ToLower(prompt(r, "회수된 메시지도 포함? (y/N)", "N")) == "y"
	if gid == "" {
		fmt.Println("GroupId는 필수입니다.")
		return
	}
	var seq uint64
	if before != "" {
		if v, _ := strconv.ParseUint(before, 10, 64); v > 0 {
			seq = v
		}
	}
	sp := NewSpinner()
	sp.Start("조회 중")
	raw, _, err := c.GetGroupHistory(ctx, gid, n, seq, withRec)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

func handleGroupRecall(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[그룹 메시지 회수]")
	gid := prompt(r, "GroupId", "")
	seqs := prompt(r, "MsgSeq(쉼표로 여러 개)", "")
	reason := prompt(r, "Reason(선택)", "")
	if gid == "" || seqs == "" {
		fmt.Println("GroupId/MsgSeq는 필수입니다.")
		return
	}
	var list []uint64
	for _, s := range strings.Split(seqs, ",") {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		if v, err := strconv.ParseUint(s, 10, 64); err == nil {
			list = append(list, v)
		}
	}
	if len(list) == 0 {
		fmt.Println("유효한 MsgSeq가 없습니다.")
		return
	}
	sp := NewSpinner()
	sp.Start("회수 중")
	raw, _, err := c.RecallGroupMsgs(ctx, gid, reason, list...)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

// ========== Handlers: Accounts ==========

func handleAccountImport(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[계정 등록(AccountImport)]")
	id := prompt(r, "UserID", "")
	nick := prompt(r, "Nick(선택)", "")
	face := prompt(r, "FaceUrl(선택)", "")
	if id == "" {
		fmt.Println("UserID는 필수입니다.")
		return
	}
	sp := NewSpinner()
	sp.Start("등록 중")
	raw, _, err := c.AccountImport(ctx, id, nick, face)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

func handleAccountCheck(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[계정 체크(AccountCheck)]")
	line := prompt(r, "UserIDs(쉼표로 구분)", "")
	if strings.TrimSpace(line) == "" {
		fmt.Println("UserIDs는 필수입니다.")
		return
	}
	var ids []string
	for _, s := range strings.Split(line, ",") {
		if v := strings.TrimSpace(s); v != "" {
			ids = append(ids, v)
		}
	}
	sp := NewSpinner()
	sp.Start("조회 중")
	raw, _, err := c.AccountCheck(ctx, ids...)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

func handleKick(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[강제 로그아웃(Kick)]")
	id := prompt(r, "UserID", "")
	if id == "" {
		fmt.Println("UserID는 필수입니다.")
		return
	}
	sp := NewSpinner()
	sp.Start("처리 중")
	raw, _, err := c.Kick(ctx, id)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

func handleQueryOnlineStatus(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[온라인 상태 조회]")
	line := prompt(r, "UserIDs(쉼표로 구분, 최대 500)", "")
	if strings.TrimSpace(line) == "" {
		fmt.Println("UserIDs는 필수입니다.")
		return
	}
	needDetail := strings.ToLower(prompt(r, "상세 플랫폼 정보 포함? (y/N)", "N")) == "y"
	needInstid := strings.ToLower(prompt(r, "Instid 포함? (y/N)", "N")) == "y"
	needCustom := strings.ToLower(prompt(r, "CustomStatus 포함? (y/N)", "N")) == "y"
	var ids []string
	for _, s := range strings.Split(line, ",") {
		if v := strings.TrimSpace(s); v != "" {
			ids = append(ids, v)
		}
	}
	sp := NewSpinner()
	sp.Start("조회 중")
	raw, _, err := c.QueryOnlineStatus(ctx, ids, needDetail, needInstid, needCustom)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}
