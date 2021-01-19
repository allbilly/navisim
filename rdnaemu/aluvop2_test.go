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

	It("should run V_CNDMASK_B32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOP2
		state.inst.Opcode = 1

		sp := state.Scratchpad().AsVOP2()
		sp.VCC = 1
		sp.SRC0[0] = 1
		sp.SRC0[1] = 2
		sp.SRC1[0] = 3
		sp.SRC1[1] = 4
		sp.EXEC = 3

		alu.Run(state)

		Expect(sp.DST[0]).To(Equal(uint64(3)))
		Expect(sp.DST[1]).To(Equal(uint64(2)))
		Expect(sp.DST[1]).To(Equal(uint64(2)))
	})

	It("should run V_FMAC_F32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOP2
		state.inst.Opcode = 43

		sp := state.Scratchpad().AsVOP2()
		sp.SRC0[0] = uint64(float32ToBits(4))
		sp.SRC1[0] = uint64(float32ToBits(16))
		sp.DST[0] = uint64(float32ToBits(1024))
		sp.EXEC = 1

		alu.Run(state)

		Expect(asFloat32(uint32(sp.DST[0]))).To(Equal(float32(1024.0 + 16.0*4.0)))
	})

	It("should run V_MUL_F32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOP2
		state.inst.Opcode = 8

		sp := state.Scratchpad().AsVOP2()
		sp.SRC0[0] = uint64(math.Float32bits(2.0))
		sp.SRC1[0] = uint64(math.Float32bits(3.1))
		sp.EXEC = 0x1

		alu.Run(state)

		Expect(sp.DST[0]).To(Equal(uint64(math.Float32bits(float32(6.2)))))
	})

})
