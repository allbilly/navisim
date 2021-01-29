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

	It("should decode 7E040280", func() {
		buf := []byte{0x80, 0x02, 0x04, 0x7E}

		inst, err := disassembler.Decode(buf)

		Expect(err).To(BeNil())
		Expect(inst.String(nil)).
			To(Equal("v_mov_b32_e32 v2, 0"))
	})

	It("should decode D5280001 00010005", func() {
		buf := []byte{0x01, 0x00, 0x28, 0xD5, 0x05, 0x00, 0x01, 0x00}

		inst, err := disassembler.Decode(buf)

		Expect(err).To(BeNil())
		Expect(inst.String(nil)).
			To(Equal("v_add_co_ci_u32_e64 v1, s0, s5, 0, s0"))
	})

	It("should decode 38060205", func() {
		buf := []byte{0x05, 0x02, 0x06, 0x38}

		inst, err := disassembler.Decode(buf)

		Expect(err).To(BeNil())
		Expect(inst.String(nil)).
			To(Equal("v_or_b32_e32 v3, s5, v1"))
	})

	It("should decode D4040000 00020B0F", func() {
		buf := []byte{0x00, 0x00, 0x04, 0xD4, 0x0F, 0x0B, 0x02, 0x00}

		inst, err := disassembler.Decode(buf)

		Expect(err).To(BeNil())
		Expect(inst.String(nil)).
			To(Equal("v_cmp_gt_f32_e64 s0, v15, v5"))
	})

	It("should decode D4860001 0002090B", func() {
		buf := []byte{0x01, 0x00, 0x86, 0xD4, 0x0B, 0x09, 0x02, 0x00}

		inst, err := disassembler.Decode(buf)

		Expect(err).To(BeNil())
		Expect(inst.String(nil)).
			To(Equal("v_cmp_ge_i32_e64 s1, v11, v4"))
	})
})
