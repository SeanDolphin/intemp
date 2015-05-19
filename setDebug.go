package intemp

import (
	"github.com/cogger/stash"
	"golang.org/x/net/context"
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
