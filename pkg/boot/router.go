package boot

import "github.com/rwcoding/goback/pkg/api"

func NewRouter() *Router {
	return &Router{routes: map[string]Handler{}}
}

type Logic interface {
	Run() *api.Response
}

type Handler func(ctx *Context) Logic

type Router struct {
	routes map[string]Handler
}

func (r *Router) Add(route string, handler Handler) {
	r.routes[route] = handler
}

func (r *Router) Find(route string) Handler {
	if h, ok := r.routes[route]; ok {
		return h
	}
	return nil
}

// 接口 => 名称
var authority map[string]string = map[string]string{}

func AddAuthority(route, name string) {
	authority[route] = name
}

func GetAuthorities() map[string]string {
	return authority
}
