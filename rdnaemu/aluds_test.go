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
		alu.lds = make([]byte, 4096)

		state = new(mockInstState)
		state.scratchpad = make([]byte, 4096)
	})

	It("should run DS_WRITE_B32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.DS
		state.inst.Opcode = 13
		state.inst.Offset0 = 0

		sp := state.scratchpad.AsDS()
		sp.EXEC = 0x01
		sp.ADDR[0] = 100
		sp.DATA[0] = 1

		alu.Run(state)

		lds := alu.LDS()
		Expect(rdnainsts.BytesToUint32(lds[100:])).To(Equal(uint32(1)))
	})
})
