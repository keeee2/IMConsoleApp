package im

import (
	"bytes"
	"context"
	crand "crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"
)

type Client struct {
	cfg  Config
	http *http.Client

	mu           sync.RWMutex
	sig          string
	sigCreatedAt time.Time
}

func NewClient(cfg Config) *Client {
	return &Client{cfg: cfg, http: &http.Client{Timeout: cfg.HTTPTimeout}}
}

func (c *Client) randUint32() (uint32, error) {
	var n uint32
	if err := binary.Read(crand.Reader, binary.LittleEndian, &n); err != nil {
		return 0, err
	}
	return n, nil
}

func (c *Client) delayBeforeRequest() {
	if !c.cfg.RequestDelayEnabled {
		return
	}
	if c.cfg.RequestDelayMax <= c.cfg.RequestDelayMin {
		time.Sleep(c.cfg.RequestDelayMin)
		return
	}
	n, err := c.randUint32()
	if err != nil {
		time.Sleep(c.cfg.RequestDelayMin)
		return
	}
	span := c.cfg.RequestDelayMax - c.cfg.RequestDelayMin
	offset := time.Duration(uint64(n) % uint64(span))
	time.Sleep(c.cfg.RequestDelayMin + offset)
}

func (c *Client) getUserSig() (string, error) {
	const refreshBefore = 5 * time.Minute
	c.mu.RLock()
	sig := c.sig
	created := c.sigCreatedAt
	c.mu.RUnlock()
	if sig != "" && time.Since(created) < time.Duration(c.cfg.UserSigExpireSeconds)*time.Second-refreshBefore {
		return sig, nil
	}
	newSig, err := tencentyun.GenUserSig(c.cfg.SDKAppID, c.cfg.AppKey, c.cfg.AdminIdentifier, c.cfg.UserSigExpireSeconds)
	if err != nil {
		return "", fmt.Errorf("usersig 생성 실패: %w", err)
	}
	c.mu.Lock()
	c.sig = newSig
	c.sigCreatedAt = time.Now()
	c.mu.Unlock()
	return newSig, nil
}

func (c *Client) buildURL(apiPath string) (string, error) {
	sig, err := c.getUserSig()
	if err != nil {
		return "", err
	}
	r, err := c.randUint32()
	if err != nil {
		return "", fmt.Errorf("random 생성 실패: %w", err)
	}
	base := c.cfg.Domain + apiPath
	q := url.Values{}
	q.Set("sdkappid", fmt.Sprintf("%d", c.cfg.SDKAppID))
	q.Set("identifier", c.cfg.AdminIdentifier)
	q.Set("usersig", sig)
	q.Set("random", fmt.Sprintf("%d", r))
	q.Set("contenttype", "json")
	return base + "?" + q.Encode(), nil
}

func (c *Client) postJSON(ctx context.Context, apiPath string, body any) ([]byte, int, error) {
	raw, err := json.Marshal(body)
	if err != nil {
		return nil, 0, fmt.Errorf("요청 마샬 실패: %w", err)
	}

	urlStr, err := c.buildURL(apiPath)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, urlStr, bytes.NewReader(raw))
	if err != nil {
		return nil, 0, fmt.Errorf("요청 생성 실패: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	c.delayBeforeRequest()
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return b, resp.StatusCode, nil
}
