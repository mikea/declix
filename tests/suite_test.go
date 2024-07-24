package tests

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFunctionalTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FunctionalTests Suite")
}
