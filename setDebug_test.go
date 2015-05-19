package intemp_test

import (
	"fmt"
	"os"
	"path/filepath"

	. "github.com/SeanDolphin/intemp"
	"golang.org/x/net/context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SetDebug", func() {

	It("should delete the directory when debug is not set", func() {
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

	It("should delete the directory when debug is set to false", func() {
		var fileName string
		ctx := SetDebug(context.Background(), false)
		err := New(ctx, func(ctx context.Context, tempDir string) error {
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

	It("should keep the directory when debug is set to true", func() {
		var fileName string
		ctx := SetDebug(context.Background(), true)
		err := New(ctx, func(ctx context.Context, tempDir string) error {
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
		Expect(err).ToNot(HaveOccurred())
		f.Close()
	})
})
