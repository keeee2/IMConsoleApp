package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	im "tencent/api/internal"
)

// ========== Console Helpers ==========

type Session struct {
	Limit int
	Next  uint64
}

func NewSession() *Session { return &Session{Limit: 100, Next: 0} }

func banner(cfg im.Config, sess *Session) {
	fmt.Println("==============================")
	fmt.Println("       Tencent IM 콘솔        ")
	fmt.Println("==============================")
	fmt.Printf("Owner=%s, Type=%s, Limit=%d, Next=%d\n\n", cfg.AdminIdentifier, cfg.DefaultGroupType, sess.Limit, sess.Next)
}

// 상위 메뉴
func mainMenu() {
	fmt.Println("------------------------------")
	fmt.Println("무엇을 하시겠어요?")
	fmt.Println("  1) 그룹 관리")
	fmt.Println("  2) 메시지 관리")
	fmt.Println("  3) 계정 관리")
	fmt.Println("  q) 종료")
}

// 그룹 관리 하위 메뉴
func groupMenu() {
	fmt.Println("------------------------------")
	fmt.Println("[그룹 관리]")
	fmt.Println("  1) 그룹 만들기")
	fmt.Println("  2) 그룹 삭제하기")
	fmt.Println("  3) 그룹 목록 보기")
	fmt.Println("  4) 그룹 상세 조회")
	fmt.Println("  5) 멤버 목록 조회")
	fmt.Println("  6) 멤버 초대")
	fmt.Println("  7) 멤버 제거")
	fmt.Println("  8) 멤버 음소거/해제")
	fmt.Println("  b) 뒤로가기")
	fmt.Println("  q) 종료")
}

// 메시지 전송 하위 메뉴
func messageMenu() {
	fmt.Println("------------------------------")
	fmt.Println("[메시지 전송]")
	fmt.Println("  1) 시스템 메시지 전송")
	fmt.Println("  2) 그룹 일반 텍스트 전송")
	fmt.Println("  3) 1:1(C2C) 텍스트 전송")
	fmt.Println("  4) 그룹 히스토리 조회")
	fmt.Println("  5) 그룹 메시지 회수")
	fmt.Println("  b) 뒤로가기")
	fmt.Println("  q) 종료")
}

// 계정 관리 하위 메뉴
func accountMenu() {
	fmt.Println("------------------------------")
	fmt.Println("[계정 관리]")
	fmt.Println("  1) 계정 등록(AccountImport)")
	fmt.Println("  2) 계정 체크(AccountCheck)")
	fmt.Println("  3) 강제 로그아웃(Kick)")
	fmt.Println("  4) 온라인 상태 조회(QueryOnlineStatus)")
	fmt.Println("  b) 뒤로가기")
	fmt.Println("  q) 종료")
}

func prompt(r *bufio.Reader, label, def string) string {
	if def != "" {
		fmt.Printf("%s [%s]: ", label, def)
	} else {
		fmt.Printf("%s: ", label)
	}
	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	if line == "" {
		return def
	}
	return line
}
func promptInt(r *bufio.Reader, label string, def int) int {
	for {
		ans := prompt(r, label, fmt.Sprintf("%d", def))
		v, err := strconv.Atoi(ans)
		if err == nil {
			return v
		}
		fmt.Println("숫자를 입력해주세요.")
	}
}
func promptUint64(r *bufio.Reader, label string, def uint64) uint64 {
	for {
		ans := prompt(r, label, fmt.Sprintf("%d", def))
		v, err := strconv.ParseUint(ans, 10, 64)
		if err == nil {
			return v
		}
		fmt.Println("정수를 입력해주세요.")
	}
}

// ========== Submenu Runners & Entry ==========

func runGroupMenu(ctx context.Context, c *im.Client, r *bufio.Reader, cfg im.Config, sess *Session) {
	for {
		groupMenu()
		choice := prompt(r, "\n번호를 입력하세요", "1 ~ 8")
		fmt.Println()
		switch strings.ToLower(choice) {
		case "1":
			handleCreate(ctx, c, r, cfg)
		case "2":
			handleDestroy(ctx, c, r)
		case "3":
			handleList(ctx, c, r, sess)
		case "4":
			handleGetGroupInfo(ctx, c, r)
		case "5":
			handleGetGroupMembers(ctx, c, r)
		case "6":
			handleAddMembers(ctx, c, r)
		case "7":
			handleDeleteMembers(ctx, c, r)
		case "8":
			handleMuteMembers(ctx, c, r)
		case "b":
			return
		case "q", "quit", "exit":
			fmt.Println("\n종료합니다... 👋")
			os.Exit(0)
		default:
			fmt.Println("\n알 수 없는 선택입니다. 1 ~ 8, b, q 중에서 골라주세요.")
		}
	}
}

func runMessageMenu(ctx context.Context, c *im.Client, r *bufio.Reader) {
	for {
		messageMenu()
		choice := prompt(r, "\n번호를 입력하세요", "1 ~ 5")
		fmt.Println()
		switch strings.ToLower(choice) {
		case "1":
			handleSendSystemMessage(ctx, c, r)
		case "2":
			handleSendGroupText(ctx, c, r)
		case "3":
			handleSendC2CText(ctx, c, r)
		case "4":
			handleGroupHistory(ctx, c, r)
		case "5":
			handleGroupRecall(ctx, c, r)
		case "b":
			return
		case "q", "quit", "exit":
			fmt.Println("\n종료합니다... 👋")
			os.Exit(0)
		default:
			fmt.Println("\n알 수 없는 선택입니다. 1 ~ 5, b, q 중에서 골라주세요.")
		}
	}
}

func runAccountMenu(ctx context.Context, c *im.Client, r *bufio.Reader) {
	for {
		accountMenu()
		choice := prompt(r, "\n번호를 입력하세요", "1 ~ 4")
		fmt.Println()
		switch strings.ToLower(choice) {
		case "1":
			handleAccountImport(ctx, c, r)
		case "2":
			handleAccountCheck(ctx, c, r)
		case "3":
			handleKick(ctx, c, r)
		case "4":
			handleQueryOnlineStatus(ctx, c, r)
		case "b":
			return
		case "q", "quit", "exit":
			fmt.Println("\n종료합니다... 👋")
			os.Exit(0)
		default:
			fmt.Println("\n알 수 없는 선택입니다. 1 ~ 4, b, q 중에서 골라주세요.")
		}
	}
}

func StartConsole(client *im.Client, cfg im.Config) {
	r := bufio.NewReader(os.Stdin)
	sess := NewSession()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go func() {
		<-ctx.Done()
		fmt.Fprint(os.Stderr, "\r\x1b[2K\x1b[?25h\n")
		fmt.Println("종료합니다... 👋")
		os.Exit(0)
	}()

	banner(cfg, sess)

	for {
		mainMenu()
		choiceStr := prompt(r, "\n번호를 입력하세요", "1 ~ 3")
		fmt.Println()
		switch strings.ToLower(choiceStr) {
		case "1":
			runGroupMenu(ctx, client, r, cfg, sess)
		case "2":
			runMessageMenu(ctx, client, r)
		case "3":
			runAccountMenu(ctx, client, r)
		case "q", "quit", "exit":
			fmt.Println("\n종료합니다... 👋")
			return
		default:
			fmt.Println("\n알 수 없는 선택입니다. 1 ~ 3 중에서 골라주세요.")
		}
	}
}
