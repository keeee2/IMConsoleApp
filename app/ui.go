package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	im "tencent/api/internal"
)

// ========== UX: Spinner + Emoji Status ==========

type Spinner struct {
	msg    string
	frames []rune
	ticker *time.Ticker
	quit   chan struct{}
	done   chan struct{}
	active bool
}

func NewSpinner() *Spinner {
	return &Spinner{frames: []rune{'⠋', '⠙', '⠹', '⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏'}}
}

func (s *Spinner) Start(msg string) {
	if s.active {
		return
	}
	s.msg = msg
	s.quit = make(chan struct{})
	s.done = make(chan struct{})
	s.ticker = time.NewTicker(90 * time.Millisecond)
	s.active = true
	fmt.Fprint(os.Stderr, "\x1b[?25l")
	go func() {
		defer close(s.done)
		idx := 0
		for {
			select {
			case <-s.quit:
				return
			case <-s.ticker.C:
				fmt.Fprint(os.Stderr, "\r\x1b[2K")
				fmt.Fprintf(os.Stderr, "%s %c", s.msg, s.frames[idx%len(s.frames)])
				idx++
			}
		}
	}()
}

func (s *Spinner) Stopf(format string, a ...any) {
	if !s.active {
		return
	}
	s.active = false
	s.ticker.Stop()
	close(s.quit)
	<-s.done
	fmt.Fprint(os.Stderr, "\r\x1b[2K")
	if format != "" {
		fmt.Fprintf(os.Stderr, "%s\n", fmt.Sprintf(format, a...))
	} else {
		fmt.Fprintln(os.Stderr)
	}
	fmt.Fprint(os.Stderr, "\x1b[?25h")
}

const (
	iconOK   = "✅"
	iconFail = "❌"
	iconWarn = "⚠️"
)

func printAPISummary(res *im.ResponseData, successMsg string) {
	if res == nil {
		fmt.Printf("%s 응답이 비어있어요. 원본 JSON을 확인하세요.\n", iconWarn)
		return
	}
	if res.ErrorCode != 0 {
		fmt.Printf("%s 실패 (ErrorCode=%d, Info=%s)\n", iconFail, res.ErrorCode, res.ErrorInfo)
		return
	}
	fmt.Printf("%s %s\n", iconOK, successMsg)
}

func prettyPrintJSON(raw []byte) {
	var buf bytes.Buffer
	if err := json.Indent(&buf, raw, "", "  "); err != nil {
		fmt.Println(string(raw))
		return
	}
	fmt.Println(buf.String())
}
