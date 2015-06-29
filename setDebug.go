package intemp

import (
	"golang.org/x/net/context"
	"gopkg.in/cogger/stash.v1"
)

func SetDebug(ctx context.Context, flag bool) context.Context {
	return stash.Set(ctx, "isDebug", flag)
}

func isDebug(ctx context.Context) bool {
	debug, _ := stash.Get(ctx, "isDebug", func() interface{} {
		return false
	}).(bool)
	return debug
}
