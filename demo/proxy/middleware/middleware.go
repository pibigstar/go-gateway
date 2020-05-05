package middleware

import (
	"context"
	"math"
	"net/http"
	"strings"
)

// 中间件核心结构体

//最多 63 个中间件
const abortIndex = math.MaxInt8 / 2

type HandlerFunc func(*RouterContext)

// 路由
type Router struct {
	groups []*RouterGroup
}

// 每个path下对应得路由组
type RouterGroup struct {
	*Router
	path     string
	Handlers []HandlerFunc
}

// 路由上下文
type RouterContext struct {
	Rw    http.ResponseWriter
	Req   *http.Request
	Ctx   context.Context
	Index int // 当前要执行得中间件索引值
	*RouterGroup
}

// 路由控制器
type RouterHandler struct {
	core   func(*RouterContext) http.Handler
	router *Router
}

// 实现Handler接口
func (w *RouterHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	c := newRouterContext(rw, req, w.router)
	if w.core != nil {
		c.Handlers = append(c.Handlers, func(c *RouterContext) {
			w.core(c).ServeHTTP(rw, req)
		})
	}
	c.Reset()
	c.Next()
}

// 创建Router上下文
func newRouterContext(rw http.ResponseWriter, req *http.Request, r *Router) *RouterContext {
	routerGroup := &RouterGroup{}
	//最长url前缀匹配
	matchUrlLen := 0
	for _, group := range r.groups {
		if strings.HasPrefix(req.RequestURI, group.path) {
			pathLen := len(group.path)
			if pathLen > matchUrlLen {
				matchUrlLen = pathLen
				//浅拷贝数组指针
				*routerGroup = *group
			}
		}
	}

	c := &RouterContext{Rw: rw, Req: req, RouterGroup: routerGroup, Ctx: req.Context()}
	c.Reset()
	return c
}

func (c *RouterContext) Get(key interface{}) interface{} {
	return c.Ctx.Value(key)
}

func (c *RouterContext) Set(key, val interface{}) {
	c.Ctx = context.WithValue(c.Ctx, key, val)
}

// 执行下一个Handler中间件
func (c *RouterContext) Next() {
	c.Index++
	for c.Index < len(c.Handlers) {
		c.Handlers[c.Index](c)
		c.Index++
	}
}

// 跳出中间件方法
func (c *RouterContext) Abort() {
	c.Index = abortIndex
}

// 是否跳过了回调
func (c *RouterContext) IsAborted() bool {
	return c.Index >= abortIndex
}

// 重置回调
func (c *RouterContext) Reset() {
	c.Index = -1
}

func NewRouterHandler(coreFunc func(*RouterContext) http.Handler, router *Router) *RouterHandler {
	return &RouterHandler{
		core:   coreFunc,
		router: router,
	}
}

// 构造 router
func NewRouter() *Router {
	return &Router{}
}

// 创建 Group
func (g *Router) Group(path string) *RouterGroup {
	return &RouterGroup{
		Router: g,
		path:   path,
	}
}

// 构造回调方法
func (g *RouterGroup) Use(meddlers ...HandlerFunc) *RouterGroup {
	g.Handlers = append(g.Handlers, meddlers...)
	existsFlag := false
	for _, oldGroup := range g.Router.groups {
		if oldGroup == g {
			existsFlag = true
		}
	}
	if !existsFlag {
		g.Router.groups = append(g.Router.groups, g)
	}
	return g
}
