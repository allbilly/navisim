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

	It("should run V_CVT_F32_U32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOP1
		state.inst.Opcode = 6

		sp := state.Scratchpad().AsVOP1()
		sp.SRC0[0] = 64
		sp.EXEC = 0x1

		alu.Run(state)

		Expect(math.Float32frombits(uint32(sp.DST[0]))).To(Equal(float32(64.0)))
	})

	It("should run V_CVT_U32_F32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOP1
		state.inst.Opcode = 7

		sp := state.Scratchpad().AsVOP1()
		sp.SRC0[0] = uint64(math.Float32bits(64))
		sp.EXEC = 0x1

		alu.Run(state)
		Expect(sp.DST[0]).To(Equal(uint64(64)))
	})

	It("should run V_RCP_IFLAG_F32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOP1
		state.inst.Opcode = 43

		sp := state.Scratchpad().AsVOP1()
		sp.SRC0[0] = uint64(math.Float32bits(4.0))
		sp.EXEC = 0x3

		alu.Run(state)
		Expect(math.Float32frombits(uint32(sp.DST[0]))).To(Equal(float32(0.25)))
	})
})
