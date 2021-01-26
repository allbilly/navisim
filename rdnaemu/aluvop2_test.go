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
	It("should run v_max_f32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOP2
		state.inst.Opcode = 16

		sp := state.Scratchpad().AsVOP2()
		sp.SRC0[0] = uint64(float32ToBits(4))
		sp.SRC1[0] = uint64(float32ToBits(16))
		sp.EXEC = 1

		alu.Run(state)

		Expect(asFloat32(uint32(sp.DST[0]))).To(Equal(float32(16)))
	})
	It("should run v_add_co_ci", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOP2
		state.inst.Opcode = 40

		sp := state.Scratchpad().AsVOP2()
		sp.SRC0[0] = 0xf0000000
		sp.SRC1[0] = 0x0000f000
		sp.SRC0[1] = 0xf0000000
		sp.SRC1[1] = 0x1000f000
		sp.VCC = 0x00000001
		sp.EXEC = 3

		alu.Run(state)

		Expect(sp.VCC).To(Equal(uint64(0x00000002)))
		Expect(sp.DST[0]).To(Equal(uint64(0xf000f001)))
		Expect(sp.DST[1]).To(Equal(uint64(0x0000f000)))
	})

	It("should run v_add_nc_u32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOP2
		state.inst.Opcode = 37

		sp := state.Scratchpad().AsVOP2()
		sp.SRC0[0] = math.MaxUint32 - 10
		sp.SRC1[0] = 11
		sp.VCC = 0
		sp.EXEC = 1

		alu.Run(state)

		Expect(sp.VCC).To(Equal(uint64(0)))
		Expect(sp.DST[0]).To(Equal(uint64(0)))
	})

	It("should run v_or_b32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOP2
		state.inst.Opcode = 28

		sp := state.Scratchpad().AsVOP2()
		sp.SRC0[0] = 0x0000ffff
		sp.SRC1[0] = 0xff000000
		sp.EXEC = 1

		alu.Run(state)

		Expect(sp.DST[0]).To(Equal(uint64(0xff00ffff)))
	})
})
