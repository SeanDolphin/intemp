package intemp

import (
	"github.com/cogger/stash"
	"golang.org/x/net/context"
)

func SetRoot(ctx context.Context, path string) context.Context {
	return stash.Set(ctx, "prefix", path)
}

func getRoot(ctx context.Context) string {
	prefix, _ := stash.Get(ctx, "prefix", func() interface{} {
		return ""
	}).(string)
	return prefix
}
