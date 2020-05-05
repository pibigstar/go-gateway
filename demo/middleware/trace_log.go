package middleware

import "log"

// 打印trace中间件
func TraceLogSliceMW() func(c *RouterContext) {
	return func(c *RouterContext) {
		log.Println("trace_in")
		c.Next()
		log.Println("trace_out")
	}
}
