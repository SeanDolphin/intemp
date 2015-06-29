package intemp

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/net/context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Intemp", func() {
	It("should return any errors from inside the temp scope", func() {
		var ErrSomeRandom = errors.New("some random error")
		err := New(context.Background(), func(ctx context.Context, tempDir string) error {
			return ErrSomeRandom
		})
		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal(ErrSomeRandom))
	})

	It("should delete the directory when it exits", func() {
		var fileName string
		err := New(context.Background(), func(ctx context.Context, tempDir string) error {
			fileName = filepath.Join(tempDir, "main.go")
			f, err := os.Create(fileName)
			if err != nil {
				return err
			}

			defer f.Close()

			fmt.Fprint(f, "BLAH")
			return nil
		})

		Expect(err).ToNot(HaveOccurred())

		f, err := os.Open(fileName)
		Expect(err).To(HaveOccurred())
		f.Close()
	})

	It("should error when a temp directory can not be created", func() {
		ctx := SetRoot(context.Background(), "/_does_not_exist")
		err := New(ctx, func(ctx context.Context, tempDir string) error {
			return nil
		})

		Expect(err).To(HaveOccurred())
	})
})
