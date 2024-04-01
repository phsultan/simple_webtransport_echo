package main

import (
  "log"
	"os"

	"github.com/joho/godotenv"

  "github.com/phsultan/simple_webtransport_echo/internal"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("%s, taking environment variables from shell", err)
	}

	CERT_FILE := os.Getenv("CERT_FILE")
	KEY_FILE := os.Getenv("KEY_FILE")
	PORT := "4443"
	PATH := "/webtransport"

	if CERT_FILE == "" {
		log.Fatal("CERT_FILE environment variable not set")
	}

	if KEY_FILE == "" {
		log.Fatal("KEY_FILE environment variable not set")
	}

  log.Printf("Start listening on port %s and path %s", PORT, PATH)
  err = webtransport.StartServer(CERT_FILE, KEY_FILE, PORT, PATH)
  log.Printf("Stopped, err : %s", err)
}
