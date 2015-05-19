package intemp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIntemp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Intemp Suite")
}
