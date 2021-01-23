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

	It("should run v_cmp_gt_i32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOPC
		state.inst.Opcode = 132

		sp := state.Scratchpad().AsVOPC()
		sp.EXEC = 0xF
		sp.SRC0[0] = 1
		sp.SRC0[1] = uint64(int32ToBits(-1))
		sp.SRC0[2] = 1
		sp.SRC0[3] = 1
		sp.SRC1[0] = 1
		sp.SRC1[1] = uint64(int32ToBits(-2))
		sp.SRC1[2] = 0
		sp.SRC1[3] = 2

		alu.Run(state)

		Expect(sp.VCC).To(Equal(uint64(0x6)))
	})

	It("should run v_cmp_gt_u32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOPC
		state.inst.Opcode = 196

		sp := state.Scratchpad().AsVOPC()
		sp.EXEC = 0x7
		sp.SRC0[0] = 1
		sp.SRC1[0] = 1
		sp.SRC0[1] = 1
		sp.SRC1[1] = 2
		sp.SRC0[2] = 1
		sp.SRC1[2] = 0

		alu.Run(state)

		Expect(sp.VCC).To(Equal(uint64(0x4)))
	})

})
