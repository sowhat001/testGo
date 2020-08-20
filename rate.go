package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
	"time"
)

func TestRate() {
	limiter := rate.NewLimiter(5, 10)
	count := 0
	for {
		count++
		if count > 100 {
			break
		}

		err := limiter.WaitN(context.Background(), 100)
		if err != nil {
			logrus.Error(count, "no")
		}
		time.Sleep(1000)
	}
}
