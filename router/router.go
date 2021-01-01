package router

import (
	"edboffical/webdog/config"
	"edboffical/webdog/context"
	"edboffical/webdog/manager"
	"edboffical/webdog/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type DogRouter struct {
	router *httprouter.Router
}

// init server router
func InitDogRouter() *DogRouter {
	dr := &DogRouter{}
	dr.router = httprouter.New()

	dr.RegisterRouter()

	return dr
}

// registers a route that matches all the HTTP methods.
// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
func (wr *DogRouter) Any(relativePath string, handler httprouter.Handle) {
	wr.router.Handle(http.MethodGet, relativePath, handler)
	wr.router.Handle(http.MethodPost, relativePath, handler)
	wr.router.Handle(http.MethodPut, relativePath, handler)
	wr.router.Handle(http.MethodPatch, relativePath, handler)
	wr.router.Handle(http.MethodHead, relativePath, handler)
	wr.router.Handle(http.MethodOptions, relativePath, handler)
	wr.router.Handle(http.MethodDelete, relativePath, handler)
	wr.router.Handle(http.MethodConnect, relativePath, handler)
	wr.router.Handle(http.MethodTrace, relativePath, handler)
}

// registers a route that matches assigned method
func (wr *DogRouter) SetHandler(method, relativePath string, handler httprouter.Handle) {
	wr.router.Handle(method, relativePath, handler)
}

// return router
func (wr *DogRouter) GetRouter() *httprouter.Router {
	return wr.router
}

// register or update router
func (dr *DogRouter) RegisterRouter() {
	cfg := config.ReadCfg()
	// register default router
	for router, config := range cfg.Normal {
		var ctx *context.Context

		// if registered
		var handle httprouter.Handle
		if config.Method == "" {
			handle, _, _ = dr.router.Lookup("GET", router)
		} else {
			handle, _, _ = dr.router.Lookup(config.Method, router)
		}
		if handle != nil {
			// get ctx from manager
			mr := utils.GetMr(config.Method, router)
			ctx = manager.Get(mr)

			ctx.SetCfg(config)
			continue
		}

		ctx = context.GetCtx(config)
		var handler httprouter.Handle
		switch config.Mode {
		case "command":
			handler = ctx.CommandHandler
		case "content":
			handler = ctx.ContentHandler
		default:
			handler = ctx.CommonHandler
		}

		if config.Method == "" {
			dr.Any(router, handler)
		} else {
			dr.SetHandler(config.Method, router, handler)
		}
		// store ctx
		mr := utils.GetMr(config.Method, router)
		manager.Store(mr, ctx)
	}
}
