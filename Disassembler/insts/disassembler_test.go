package insts

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDisassembler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RDNA Disassembler")
}

var _ = Describe("Disassembler", func() {
	var (
		disassembler *Disassembler
	)

	BeforeEach(func() {
		disassembler = NewDisassembler()
	})

	It("should disassemble DC308000 037D0005", func() {
		buf := []byte{0x00, 0x80, 0x30, 0xDC, 0x05, 0x00, 0x7D, 0x03}

		inst, err := disassembler.Decode(buf)

		Expect(err).To(BeNil())
		Expect(inst.String(nil)).To(Equal("global_load_dword v3, v[5:6], off"))
	})
})
