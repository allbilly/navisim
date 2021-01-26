package rdnaemu

import (
	"math"

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

	It("should run V_CVT_F32_U32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOP1
		state.inst.Opcode = 6

		sp := state.Scratchpad().AsVOP1()
		sp.SRC0[0] = 1
		sp.EXEC = 0x1

		alu.Run(state)

		Expect(math.Float32frombits(uint32(sp.DST[0]))).To(Equal(float32(1.0)))
	})
})
