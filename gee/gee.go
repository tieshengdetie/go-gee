package gee

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router *router
}

// New engine 的构造方法
func New() *Engine {

	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := NewContext(writer, request)
	engine.router.handle(c)
}

func (engine *Engine) Run(addr string) (err error) {

	return http.ListenAndServe(addr, engine)
}
