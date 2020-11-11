package main

import (
	"context"
	"fmt"
	"time"
)

func TestDeadlockCtx() {
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	go func() {
		select {
		case <-ctxWithTimeout.Done():
			fmt.Printf("recv done, %v\n", ctxWithTimeout.Err())
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("ctxTest return, %v\n", ctx.Err())
		return
	}
}
