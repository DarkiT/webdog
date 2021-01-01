package manager

import (
	"edboffical/webdog/context"
	"sync"
)

var (
	Ctxs sync.Map
)

// store ctx
func Store(mr string, ctx *context.Context) {
	Ctxs.Store(mr, ctx)
}

// get ctx
func Get(mr string) *context.Context {
	val, ok := Ctxs.Load(mr)
	if !ok {
		return nil
	}
	v, ok := val.(*context.Context)
	if !ok {
		return nil
	}
	return v
}
