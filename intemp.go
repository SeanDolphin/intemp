package intemp

import (
	"io/ioutil"
	"os"

	"github.com/cogger/stash"
	"golang.org/x/net/context"
)

func New(ctx context.Context, f func(context.Context, string) error) error {
	prefix, _ := stash.Get(ctx, "prefix", func() interface{} {
		return ""
	}).(string)

	temp, err := ioutil.TempDir(prefix, "intemp")
	if err != nil {
		return err
	}

	defer os.RemoveAll(temp)
	return f(ctx, temp)
}
