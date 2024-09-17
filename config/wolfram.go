package config

import (
	"os"
	"sync"

	"github.com/krognol/go-wolfram"
)

var (
	wolfClient   *wolfram.Client
	wolfInstance sync.Once
)

func GetWolframClient() *wolfram.Client {
	wolfInstance.Do(func() {
		wolfAppID := os.Getenv("WOLFRAM_APP_ID")
		if wolfAppID == "" {
			panic("Missing WOLFRAM_APP_ID")
		}

		wolfClient = &wolfram.Client{
			AppID: wolfAppID,
		}
	})

	return wolfClient
}
