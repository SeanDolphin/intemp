package intemp_test

import (
	"os"
	"strings"

	. "github.com/SeanDolphin/intemp"
	"golang.org/x/net/context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SetRoot", func() {
	It("should set the root of a temp folder", func() {
		gopath := os.Getenv("GOPATH")
		ctx := SetRoot(context.Background(), gopath)

		err := New(ctx, func(ctx context.Context, tempDir string) error {
			Expect(strings.Contains(tempDir, gopath)).To(BeTrue())
			return nil
		})

		Expect(err).ToNot(HaveOccurred())
	})
})
