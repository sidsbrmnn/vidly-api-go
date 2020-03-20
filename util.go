package main

import (
	"context"
	"time"
)

func dbContext(i time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), i*time.Second)
	return ctx, cancel
}
