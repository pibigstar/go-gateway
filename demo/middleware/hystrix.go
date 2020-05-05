package middleware

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
)

import (
	"github.com/afex/hystrix-go/hystrix"
)

// 熔断，降级
func CircuitMW() func(c *RouterContext) {
	return func(c *RouterContext) {
		err := hystrix.Do("common", func() error {
			c.Next()
			statusCode, ok := c.Get("status_code").(int)
			if !ok || statusCode != 200 {
				return errors.New("downstream error")
			}
			return nil
		}, func(err error) error {
			fmt.Println("这里做服务降级处理.....")
			return nil
		})

		if err != nil {
			//加入自动降级处理，如获取缓存数据等
			switch err {
			case hystrix.ErrCircuitOpen:
				c.Rw.Write([]byte("circuit error:" + err.Error()))
			case hystrix.ErrMaxConcurrency:
				c.Rw.Write([]byte("circuit error:" + err.Error()))
			default:
				c.Rw.Write([]byte("circuit error:" + err.Error()))
			}
			c.Abort()
		}
	}
}

// 配置熔断信息
func SetHystrixConf(openStream bool) {
	hystrix.ConfigureCommand("common", hystrix.CommandConfig{
		Timeout:                1000, // 单次请求 超时时间
		MaxConcurrentRequests:  1,    // 最大并发量
		RequestVolumeThreshold: 1,    // 10秒内如果出现1次错误就触发熔断
		ErrorPercentThreshold:  1,    // 按百分比，如果出现1%的错误就触发熔断
		SleepWindow:            5000, // 熔断后5秒后再去尝试服务是否可用
	})

	if openStream {
		hystrixStreamHandler := hystrix.NewStreamHandler()
		hystrixStreamHandler.Start()
		go func() {
			err := http.ListenAndServe(net.JoinHostPort("", "8001"), hystrixStreamHandler)
			log.Fatal(err)
		}()
	}
}
