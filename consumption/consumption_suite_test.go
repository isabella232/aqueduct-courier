package consumption_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUsage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Consumption Suite")
}
