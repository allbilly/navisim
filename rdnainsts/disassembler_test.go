package rdnainsts

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

	It("should disassembler BE803C6A", func() {
		buf := []byte{0x6A, 0x3C, 0x80, 0xBE}

		inst, err := disassembler.Decode(buf)

		Expect(err).To(BeNil())
		Expect(inst.String(nil)).To(Equal("s_and_saveexec_b32 s0, vcc_lo"))
	})

	It("should disassembler 50060201", func() {
		buf := []byte{0x01, 0x02, 0x06, 0x50}

		inst, err := disassembler.Decode(buf)

		Expect(err).To(BeNil())
		Expect(inst.String(nil)).
			To(Equal("v_add_co_ci_u32_e32 v3, vcc_lo, s1, v1, vcc_lo"))
	})

	It("should disassemble DC308000 037D0005", func() {
		buf := []byte{0x00, 0x80, 0x30, 0xDC, 0x05, 0x00, 0x7D, 0x03}

		inst, err := disassembler.Decode(buf)

		Expect(err).To(BeNil())
		Expect(inst.String(nil)).To(Equal("global_load_dword v3, v[5:6], off"))
	})

	It("should disassemble D7010000 0002029E", func() {
		buf := []byte{0x00, 0x00, 0x01, 0xD7, 0x9E, 0x02, 0x02, 0x00}

		inst, err := disassembler.Decode(buf)

		Expect(err).To(BeNil())
		Expect(inst.String(nil)).To(Equal("v_ashrrev_i64 v[0:1], 30, v[1:2]"))
	})

	It("should decode F4000082 FA000004", func() {
		buf := []byte{0x82, 0x00, 0x00, 0xF4, 0x04, 0x00, 0x00, 0xFA}

		inst, err := disassembler.Decode(buf)

		Expect(err).To(BeNil())
		Expect(inst.String(nil)).
			To(Equal("s_load_dword s2, s[4:5], 0x4"))
	})
})
