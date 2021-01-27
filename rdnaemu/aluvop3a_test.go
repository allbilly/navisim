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

	It("should run V_CNDMASK_B32 VOP3a", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOP3a
		state.inst.Opcode = 257

		sp := state.Scratchpad().AsVOP3A()
		sp.SRC0[0] = 1
		sp.SRC1[0] = 2
		sp.SRC0[1] = 1
		sp.SRC1[1] = 2
		sp.SRC2[0] = 1
		sp.EXEC = 3

		alu.Run(state)

		Expect(sp.DST[0]).To(Equal(uint64(2)))
		Expect(sp.DST[1]).To(Equal(uint64(1)))
	})

	It("should run V_LSHL_REV B64", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOP3a
		state.inst.Opcode = 767

		sp := state.Scratchpad().AsVOP3A()

		sp.SRC1[0] = uint64(0x0000000000010000)
		sp.SRC0[0] = uint64(3)
		sp.EXEC = 0x1

		alu.Run(state)

		Expect(sp.DST[0]).To(Equal(uint64(0x0000000000080000)))

	})

	It("should run V_MUL_LO_U32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOP3a
		state.inst.Opcode = 361

		sp := state.Scratchpad().AsVOP3A()
		for i := 0; i < 64; i++ {
			sp.SRC0[i] = uint64(i)
			sp.SRC1[i] = uint64(2)
		}
		sp.EXEC = 0xffffffffffffffff

		alu.Run(state)

		for i := 0; i < 64; i++ {
			Expect(sp.DST[i]).To(Equal(uint64(i * 2)))
		}
	})

	It("should run V_MUL_HI_U32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOP3a
		state.inst.Opcode = 362

		sp := state.Scratchpad().AsVOP3A()
		sp.SRC0[0] = uint64(0x80000000)
		sp.SRC1[0] = uint64(2)
		sp.EXEC = 1

		alu.Run(state)

		Expect(sp.DST[0]).To(Equal(uint64(1)))
	})

	It("should run V_LSHL_ADD_U32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOP3a
		state.inst.Opcode = 838

		sp := state.Scratchpad().AsVOP3A()
		sp.SRC0[0] = uint64(2)
		sp.SRC1[0] = uint64(2)
		sp.SRC2[0] = uint64(1)
		sp.EXEC = 1

		alu.Run(state)

		Expect(sp.DST[0]).To(Equal(uint64(9)))
	})

	It("should run V_MAD_U32_U24", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOP3a
		state.inst.Opcode = 323

		sp := state.Scratchpad().AsVOP3A()
		sp.SRC0[0] = 10
		sp.SRC1[0] = 20
		sp.SRC2[0] = 50
		sp.EXEC = 1

		alu.Run(state)

		Expect(sp.DST[0]).To(Equal(uint64(250)))
	})
	It("should run v_cmp_ge_i32 VOP3a", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.VOP3a
		state.inst.Opcode = 134

		sp := state.Scratchpad().AsVOP3A()
		sp.SRC0[0] = uint64(int32ToBits(0))
		sp.SRC1[0] = uint64(int32ToBits(math.MinInt32))
		sp.EXEC = 1

		alu.Run(state)

		Expect(sp.VCC).To(Equal((uint64(1))))
	})

})
