package middleware

import (
	"fmt"
	"golang.org/x/time/rate"
)

func RateLimiter() func(c *RouterContext) {
	// 每秒可请求一次，最多存储两个
	l := rate.NewLimiter(1, 2)
	return func(c *RouterContext) {
		if !l.Allow() {
			c.Rw.Write([]byte(fmt.Sprintf("rate limit:%v,%v", l.Limit(), l.Burst())))
			c.Abort()
			return
		}
		c.Next()
	}
}
