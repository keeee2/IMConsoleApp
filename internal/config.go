package im

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Domain               string
	SDKAppID             int
	AppKey               string
	AdminIdentifier      string
	UserSigExpireSeconds int
	DefaultGroupType     string
	HTTPTimeout          time.Duration

	RequestDelayEnabled bool
	RequestDelayMin     time.Duration
	RequestDelayMax     time.Duration
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
func atoiOr(name, v string, def int) int {
	if v == "" {
		return def
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalf("환경변수 %s 파싱 실패: %v", name, err)
	}
	return i
}
func atob(v string) bool {
	switch strings.ToLower(v) {
	case "1", "t", "true", "y", "yes":
		return true
	}
	return false
}

func LoadConfig() Config {
	cfg := Config{
		Domain:               getenv("TENCENT_IM_DOMAIN", "https://adminapisgp.im.qcloud.com"),
		SDKAppID:             atoiOr("TENCENT_IM_SDK_APP_ID", os.Getenv("TENCENT_IM_SDK_APP_ID"), 0),
		AppKey:               getenv("TENCENT_IM_APP_KEY", ""),
		AdminIdentifier:      getenv("TENCENT_IM_ADMIN_IDENTIFIER", ""),
		UserSigExpireSeconds: atoiOr("TENCENT_IM_USERSIG_EXPIRE", os.Getenv("TENCENT_IM_USERSIG_EXPIRE"), 86400*180),
		DefaultGroupType:     getenv("TENCENT_IM_DEFAULT_GROUP_TYPE", "ChatRoom"),
		HTTPTimeout:          time.Duration(atoiOr("HTTP_TIMEOUT_SEC", os.Getenv("HTTP_TIMEOUT_SEC"), 15)) * time.Second,
		RequestDelayEnabled:  atob(getenv("REQUEST_DELAY_ENABLED", "true")),
		RequestDelayMin:      time.Duration(atoiOr("REQUEST_DELAY_MIN_MS", os.Getenv("REQUEST_DELAY_MIN_MS"), 100)) * time.Millisecond,
		RequestDelayMax:      time.Duration(atoiOr("REQUEST_DELAY_MAX_MS", os.Getenv("REQUEST_DELAY_MAX_MS"), 500)) * time.Millisecond,
	}
	if cfg.SDKAppID == 0 || cfg.AppKey == "" || cfg.AdminIdentifier == "" {
		log.Fatal("SDK_APP_ID / APP_KEY / ADMIN_IDENTIFIER 는 필수입니다 (환경변수 또는 .env로 주입).")
	}
	if cfg.RequestDelayMax < cfg.RequestDelayMin {
		cfg.RequestDelayMax = cfg.RequestDelayMin
	}
	return cfg
}
