package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
)

// 每秒可接收10个请求，最大可运行20个请求
var limiter = NewIPRateLimiter(10, 20)

// IP限速器
func IPLimit() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		// 获取IP限速器
		limiter := limiter.GetLimiter(r.RemoteAddr)
		if !limiter.Allow() {
			http.Error(r.Response.Writer, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		r.Middleware.Next()
	}
}

type IPRateLimit struct {
	mu      *sync.RWMutex
	limiter map[string]*rate.Limiter
	r       rate.Limit
	b       int
}

func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimit {
	return &IPRateLimit{
		limiter: make(map[string]*rate.Limiter),
		mu:      &sync.RWMutex{},
		r:       r, // 1s创建多少个令牌
		b:       b, // 最大存储多少个令牌
	}
}

func (i *IPRateLimit) AddIp(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()
	limiter := rate.NewLimiter(i.r, i.b)
	i.limiter[ip] = limiter
	return limiter
}

func (i *IPRateLimit) GetLimiter(ip string) *rate.Limiter {
	if limiter, ok := i.limiter[ip]; ok {
		return limiter
	}
	return i.AddIp(ip)
}