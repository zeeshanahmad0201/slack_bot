package helpers

import (
	"context"
	"time"
)

// ContextWithTimeout creates with context with timeout of 10 secs by default
func ContextWithTimeout(timeout ...time.Duration) (context.Context, context.CancelFunc) {
	if len(timeout) == 0 {
		timeout = append(timeout, time.Second*10)
	}

	return context.WithTimeout(context.Background(), timeout[0])
}

func ContextWithCancel() (context.Context, context.CancelFunc) {
	return context.WithCancel(context.Background())
}
