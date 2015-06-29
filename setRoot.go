package intemp

import (
	"golang.org/x/net/context"
	"gopkg.in/cogger/stash.v1"
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
