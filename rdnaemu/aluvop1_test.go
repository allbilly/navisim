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

	It("should run V_MOV_B32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOP1
		state.inst.Opcode = 1

		sp := state.Scratchpad().AsVOP1()
		for i := 0; i < 32; i++ {
			sp.SRC0[i] = 1
		}
		sp.EXEC = 0x00000000ffffffff

		alu.Run(state)

		for i := 0; i < 32; i++ {
			Expect(sp.SRC0[i]).To(Equal(sp.DST[i]))
		}

		for i := 32; i < 64; i++ {
			Expect(sp.SRC0[i]).To(Equal(uint64(0)))
		}

	})
})
