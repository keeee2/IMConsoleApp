package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	im "tencent/api/internal"
)

// ========== Accounts ==========
func handleMultiAccountImport(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[계정 다중 등록(MultiAccountImport)]")
	line := prompt(r, "UserIDs(쉼표)", "")
	if strings.TrimSpace(line) == "" {
		fmt.Println("UserIDs는 필수입니다.")
		return
	}
	nick := prompt(r, "Nick(전체 동일, 선택)", "")
	face := prompt(r, "FaceUrl(전체 동일, 선택)", "")
	var items []im.AccountImportReq
	for _, id := range strings.Split(line, ",") {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		items = append(items, im.AccountImportReq{UserID: id, Nick: nick, FaceUrl: face})
	}
	sp := NewSpinner()
	sp.Start("등록 중")
	raw, _, err := c.MultiAccountImport(ctx, items)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

func handleAccountDelete(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[계정 삭제(AccountDelete)]")
	line := prompt(r, "UserIDs(쉼표)", "")
	if strings.TrimSpace(line) == "" {
		fmt.Println("UserIDs는 필수입니다.")
		return
	}
	var ids []string
	for _, id := range strings.Split(line, ",") {
		if v := strings.TrimSpace(id); v != "" {
			ids = append(ids, v)
		}
	}
	sp := NewSpinner()
	sp.Start("삭제 중")
	raw, _, err := c.AccountDelete(ctx, ids...)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

// ========== C2C ==========
func handleBatchC2CText(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[1:1(C2C) 텍스트 - 여러 명]")
	tos := prompt(r, "To_Account(쉼표)", "")
	text := prompt(r, "Text", "hello everyone")
	from := prompt(r, "From_Account(빈칸=관리자)", "")
	sync := strings.ToLower(prompt(r, "다른 기기에도 동기화? (Y/n)", "Y")) != "n"
	if tos == "" || text == "" {
		fmt.Println("필수값 누락")
		return
	}
	toList := []string{}
	for _, v := range strings.Split(tos, ",") {
		if s := strings.TrimSpace(v); s != "" {
			toList = append(toList, s)
		}
	}
	sp := NewSpinner()
	sp.Start("전송 중")
	raw, _, err := c.BatchSendC2CText(ctx, toList, text, from, sync)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

func handleAdminSetMsgRead(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[1:1 읽음 처리(AdminSetMsgRead)]")
	report := prompt(r, "Report_Account", "")
	peer := prompt(r, "Peer_Account", "")
	if report == "" || peer == "" {
		fmt.Println("필수값 누락")
		return
	}
	sp := NewSpinner()
	sp.Start("처리 중")
	raw, _, err := c.AdminSetMsgRead(ctx, report, peer, 0)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

func handleAdminWithdrawC2C(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[1:1 메시지 회수(AdminMsgWithdraw)]")
	from := prompt(r, "From_Account", "")
	to := prompt(r, "To_Account", "")
	key := prompt(r, "MsgKey", "")
	if from == "" || to == "" || key == "" {
		fmt.Println("필수값 누락")
		return
	}
	sp := NewSpinner()
	sp.Start("회수 중")
	raw, _, err := c.AdminMsgWithdraw(ctx, from, to, key)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

func handleGetC2CUnread(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[1:1 미읽음 개수 조회]")
	to := prompt(r, "To_Account", "")
	peers := prompt(r, "Peer_Account 목록(쉼표, 비우면 전체)", "")
	var list []string
	for _, p := range strings.Split(peers, ",") {
		if v := strings.TrimSpace(p); v != "" {
			list = append(list, v)
		}
	}
	sp := NewSpinner()
	sp.Start("조회 중")
	raw, _, err := c.GetC2CUnreadMsgNum(ctx, to, list)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

// ========== Profile ==========
func handlePortraitSet(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[프로필 설정(PortraitSet)]")
	user := prompt(r, "UserID", "")
	tag := prompt(r, "Tag(예: Tag_Profile_IM_Nick)", "")
	val := prompt(r, "Value", "")
	if user == "" || tag == "" {
		fmt.Println("필수값 누락")
		return
	}
	items := []struct {
		Tag   string
		Value interface{}
	}{{Tag: tag, Value: val}}
	sp := NewSpinner()
	sp.Start("반영 중")
	raw, _, err := c.PortraitSet(ctx, user, items)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

func handlePortraitGet(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[프로필 조회(PortraitGet)]")
	users := prompt(r, "UserIDs(쉼표)", "")
	tags := prompt(r, "TagList(쉼표, 선택)", "")
	if strings.TrimSpace(users) == "" {
		fmt.Println("UserIDs는 필수입니다.")
		return
	}
	var uList, tList []string
	for _, u := range strings.Split(users, ",") {
		if v := strings.TrimSpace(u); v != "" {
			uList = append(uList, v)
		}
	}
	for _, t := range strings.Split(tags, ",") {
		if v := strings.TrimSpace(t); v != "" {
			tList = append(tList, v)
		}
	}
	sp := NewSpinner()
	sp.Start("조회 중")
	raw, _, err := c.PortraitGet(ctx, uList, tList)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

// ========== Global Mute ==========
func handleSetNoSpeaking(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[전역 음소거 설정(SetNoSpeaking)]")
	user := prompt(r, "UserID", "")
	c2c := uint64(promptInt(r, "C2CmsgNospeakingTime(초)", 0))
	group := uint64(promptInt(r, "GroupmsgNospeakingTime(초)", 0))
	if user == "" {
		fmt.Println("UserID는 필수입니다.")
		return
	}
	sp := NewSpinner()
	sp.Start("반영 중")
	raw, _, err := c.SetNoSpeaking(ctx, user, uint32(c2c), uint32(group))
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

func handleGetNoSpeaking(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[전역 음소거 조회(GetNoSpeaking)]")
	user := prompt(r, "UserID", "")
	if user == "" {
		fmt.Println("UserID는 필수입니다.")
		return
	}
	sp := NewSpinner()
	sp.Start("조회 중")
	raw, _, err := c.GetNoSpeaking(ctx, user)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

// ========== Ops ==========
func handleGetAppInfo(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[운영 데이터 조회(GetAppInfo)]")
	fields := prompt(r, "RequestField(쉼표, 예: ChainIncrease,ActiveUserNum)", "")
	var fl []string
	for _, f := range strings.Split(fields, ",") {
		if v := strings.TrimSpace(f); v != "" {
			fl = append(fl, v)
		}
	}
	sp := NewSpinner()
	sp.Start("조회 중")
	raw, _, err := c.GetAppInfo(ctx, fl...)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

func handleGetIPList(ctx context.Context, c *im.Client, _ *bufio.Reader) {
	fmt.Println("\n[서버 IP 조회(GetIPList)]")
	sp := NewSpinner()
	sp.Start("조회 중")
	raw, _, err := c.GetIPList(ctx)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(raw)
}

// ========== Raw API Mode ==========
func handleRawAPICall(ctx context.Context, c *im.Client, r *bufio.Reader) {
	fmt.Println("\n[Raw API 호출]")
	fmt.Println("엔드포인트 경로를 입력하세요. 예) /v4/openim/sendmsg")
	path := prompt(r, "API Path", "/v4/")
	if !strings.HasPrefix(path, "/v4/") {
		fmt.Println("/v4/ 로 시작해야 합니다.")
		return
	}
	fmt.Println("요청 JSON을 붙여넣고, 단독 줄에 '.' 을 입력해 종료하세요:")
	var b strings.Builder
	for {
		line := prompt(r, "", "")
		if strings.TrimSpace(line) == "." {
			break
		}
		b.WriteString(line)
		if !strings.HasSuffix(line, "\n") {
			b.WriteString("\n")
		}
	}
	raw := []byte(strings.TrimSpace(b.String()))
	// 간단한 JSON 유효성 검사로 UX 향상
	var js any
	if err := json.Unmarshal(raw, &js); err != nil {
		fmt.Printf("잘못된 JSON: %v\n", err)
		return
	}
	sp := NewSpinner()
	sp.Start("호출 중")
	resp, _, err := c.CallRaw(ctx, path, raw)
	if err != nil {
		sp.Stopf("실패: %v", err)
	} else {
		sp.Stopf("완료")
	}
	fmt.Println("\n--- 원본 응답(JSON) ---")
	prettyPrintJSON(resp)
}
