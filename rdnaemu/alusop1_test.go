package rdnaemu

import (
	. "github.com/onsi/ginkgo"
	"gitlab.com/akita/navisim/rdnainsts"
)

var _ = Describe("ALU", func() {

	var (
		state *mockInstState
	)

	BeforeEach(func() {

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
})
