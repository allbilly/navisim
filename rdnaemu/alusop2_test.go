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

		state = new(mockInstState)
		state.scratchpad = make([]byte, 4096)
	})

	It("should run S_SUB_I32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.SOP2
		state.inst.Opcode = 3

		sp := state.Scratchpad().AsSOP2()
		sp.SRC0 = 64
		sp.SRC1 = 10

		alu.Run(state)

		Expect(sp.DST).To(Equal(uint64(54)))
		Expect(sp.SCC).To(Equal(byte(0)))
	})
})
