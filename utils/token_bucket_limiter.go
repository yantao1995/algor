package utils

import (
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

// go rate包自带令牌桶实现
func NewTokenBucketLimiter() {
	// 每 10 秒最多允许 5 次操作（即每 2 秒发一个令牌）
	limiter := rate.NewLimiter(rate.Every(2*time.Second), 5)

	for i := 1; i <= 7; i++ {
		if limiter.Allow() {
			fmt.Printf("请求 %d 允许\n", i)
		} else {
			fmt.Printf("请求 %d 被拒绝: 超出速率限制\n", i)
		}
		time.Sleep(1 * time.Second)
	}
}
