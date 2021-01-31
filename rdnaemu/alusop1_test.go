package rdnaemu

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gitlab.com/akita/navisim/rdnainsts"
)

var _ = Describe("ALU", func() {

	var (
		alu   *ALUImpl
		state *mockInstState
	)

	BeforeEach(func() {
		alu = NewALU(nil)

		state = new(mockInstState)
		state.scratchpad = make([]byte, 4096)
	})

	It("should run s_mov_b32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.SOP1
		state.inst.Opcode = 3

		sp := state.Scratchpad().AsSOP1()
		sp.SRC0 = 0x0000ffffffff0000

		alu.Run(state)

		Expect(sp.DST).To(Equal(uint64(0x0000ffffffff0000)))
	})

	It("should run s_mov_b64", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.SOP1
		state.inst.Opcode = 4

		sp := state.Scratchpad().AsSOP1()
		sp.SRC0 = 0xffffffffffff0000

		alu.Run(state)

		Expect(sp.DST).To(Equal(uint64(0xffffffffffff0000)))
	})

	It("should run s_or_saveexec_b64", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.SOP1
		state.inst.Opcode = 61

		sp := state.Scratchpad().AsSOP1()
		sp.EXEC = 0xffffffff00000000
		sp.SRC0 = 0x0000ffffffff0000

		alu.Run(state)

		Expect(sp.EXEC).To(Equal(uint64(0xffffffffffff0000)))
		Expect(sp.DST).To(Equal(uint64(0xffffffff00000000)))
		Expect(sp.SCC).To(Equal(byte(0x1)))
	})
})
