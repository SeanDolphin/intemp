package intemp

import (
	"io/ioutil"
	"os"

	"golang.org/x/net/context"
)

func New(ctx context.Context, f func(context.Context, string) error) error {
	temp, err := ioutil.TempDir(getRoot(ctx), "intemp")
	if err != nil {
		return err
	}

	if !isDebug(ctx) {
		defer os.RemoveAll(temp)
	}
	return f(ctx, temp)
}
