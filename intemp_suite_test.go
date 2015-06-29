package intemp

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIntemp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Intemp Suite")
}
