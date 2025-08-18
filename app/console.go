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
	fmt.Println("  4) 프로필 관리")
	fmt.Println("  5) 전역 음소거(Global Mute)")
	fmt.Println("  6) 운영/진단(Ops)")
	fmt.Println("  r) Raw API 호출(모든 엔드포인트)")
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
	fmt.Println("[메시지 관리]")
	fmt.Println("  1) 그룹 메시지 전송")
	fmt.Println("  2) 1:1 메시지 전송")
	fmt.Println("  3) 그룹 시스템 알림")
	fmt.Println("  4) 그룹 히스토리 조회")
	fmt.Println("  5) 그룹 메시지 회수")
	fmt.Println("  6) 1:1(C2C) 텍스트(여러 명)")
	fmt.Println("  7) 1:1 읽음 처리(AdminSetMsgRead)")
	fmt.Println("  8) 1:1 메시지 회수(AdminMsgWithdraw)")
	fmt.Println("  9) 1:1 미읽음 개수 조회")
}

// 계정 관리 하위 메뉴
func accountMenu() {
	fmt.Println("------------------------------")
	fmt.Println("[계정 관리]")
	fmt.Println("  1) 계정 등록(AccountImport)")
	fmt.Println("  2) 계정 체크(AccountCheck)")
	fmt.Println("  3) 강제 로그아웃(Kick)")
	fmt.Println("  4) 온라인 상태 조회(QueryOnlineStatus)")
	fmt.Println("  5) 계정 다중 등록(MultiAccountImport)")
	fmt.Println("  6) 계정 삭제(AccountDelete)")
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
		case "6":
			handleBatchC2CText(ctx, c, r)
		case "7":
			handleAdminSetMsgRead(ctx, c, r)
		case "8":
			handleAdminWithdrawC2C(ctx, c, r)
		case "9":
			handleGetC2CUnread(ctx, c, r)
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
		case "4":
			// 프로필 관리
			for {
				fmt.Println("------------------------------")
				fmt.Println("[프로필 관리]")
				fmt.Println("  1) 프로필 설정(PortraitSet)")
				fmt.Println("  2) 프로필 조회(PortraitGet)")
				fmt.Println("  b) 뒤로가기")
				csel := prompt(r, "\n번호를 입력하세요", "1 ~ 2")
				fmt.Println()
				csel = strings.ToLower(csel)
				if csel == "b" {
					break
				}
				switch csel {
				case "1":
					handlePortraitSet(ctx, client, r)
				case "2":
					handlePortraitGet(ctx, client, r)
				default:
					fmt.Println("알 수 없는 선택입니다.")
				}
			}
		case "5":
			// 전역 음소거
			for {
				fmt.Println("------------------------------")
				fmt.Println("[전역 음소거]")
				fmt.Println("  1) 설정(SetNoSpeaking)")
				fmt.Println("  2) 조회(GetNoSpeaking)")
				fmt.Println("  b) 뒤로가기")
				gsel := prompt(r, "\n번호를 입력하세요", "1 ~ 2")
				fmt.Println()
				gsel = strings.ToLower(gsel)
				if gsel == "b" {
					break
				}
				switch gsel {
				case "1":
					handleSetNoSpeaking(ctx, client, r)
				case "2":
					handleGetNoSpeaking(ctx, client, r)
				default:
					fmt.Println("알 수 없는 선택입니다.")
				}
			}
		case "6":
			// 운영/진단
			for {
				fmt.Println("------------------------------")
				fmt.Println("[운영/진단]")
				fmt.Println("  1) 운영 데이터 조회(GetAppInfo)")
				fmt.Println("  2) 서버 IP 조회(GetIPList)")
				fmt.Println("  b) 뒤로가기")
				osel := prompt(r, "\n번호를 입력하세요", "1 ~ 2")
				fmt.Println()
				osel = strings.ToLower(osel)
				if osel == "b" {
					break
				}
				switch osel {
				case "1":
					handleGetAppInfo(ctx, client, r)
				case "2":
					handleGetIPList(ctx, client, r)
				default:
					fmt.Println("알 수 없는 선택입니다.")
				}
			}
		case "r":
			handleRawAPICall(ctx, client, r)
		case "q", "quit", "exit":
			fmt.Println("\n종료합니다... 👋")
			return
		default:
			fmt.Println("\n알 수 없는 선택입니다. 1 ~ 3 중에서 골라주세요.")
		}
	}
}
