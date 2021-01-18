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

	It("should run S_CMP_EQ_U32 when input is not equal", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.SOPC
		state.inst.Opcode = 6

		layout := state.Scratchpad().AsSOPC()
		layout.SRC0 = 1
		layout.SRC1 = 2

		alu.Run(state)

		Expect(layout.SCC).To(Equal(byte(0)))
	})

	It("should run S_CMP_EQ_U32 when input is equal", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.SOPC
		state.inst.Opcode = 6

		layout := state.Scratchpad().AsSOPC()
		layout.SRC0 = 1
		layout.SRC1 = 1

		alu.Run(state)

		Expect(layout.SCC).To(Equal(byte(1)))
	})

	It("should run S_CMP_LG_U32 when condition holds", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.SOPC
		state.inst.Opcode = 7

		layout := state.Scratchpad().AsSOPC()
		layout.SRC0 = 1
		layout.SRC1 = 2

		alu.Run(state)

		Expect(layout.SCC).To(Equal(byte(1)))
	})

	It("should run S_CMP_LG_U32 when condition does not hold", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.SOPC
		state.inst.Opcode = 7

		layout := state.Scratchpad().AsSOPC()
		layout.SRC0 = 1
		layout.SRC1 = 1

		alu.Run(state)

		Expect(layout.SCC).To(Equal(byte(0)))
	})

})
