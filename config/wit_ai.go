package config

import (
	"os"
	"sync"

	witai "github.com/wit-ai/wit-go"
)

var (
	client   *witai.Client
	instance sync.Once
)

func GetWitClient() *witai.Client {
	instance.Do(func() {
		witServerToken := os.Getenv("WIT_SERVER_TOKEN")
		if witServerToken == "" {
			panic("Missing WIT_SERVER_TOKEN")
		}
		client = witai.NewClient(witServerToken)
	})

	return client
}
