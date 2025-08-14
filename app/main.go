package main

import (
	"log"
	"os"

	im "tencent/api/internal"

	"github.com/joho/godotenv"
)

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	if err := godotenv.Load(".env"); err != nil {
		log.Printf("INFO: .env 로드를 건너뜀: %v", err)
	}

	cfg := im.LoadConfig()
	client := im.NewClient(cfg)

	StartConsole(client, cfg)
}
