package navisim_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNavisim(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Navisim Suite")
}
