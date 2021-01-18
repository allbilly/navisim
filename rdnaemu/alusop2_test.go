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

	It("should run S_ADD_U32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.SOP2
		state.inst.Opcode = 0

		sp := state.scratchpad.AsSOP2()
		sp.SRC0 = 1<<32 - 9
		sp.SRC1 = 10

		alu.Run(state)

		Expect(sp.DST).To(Equal(uint64(1)))
		Expect(sp.SCC).To(Equal(uint8(1)))
	})

	It("should run S_ADD_U32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.SOP2
		state.inst.Opcode = 0

		sp := state.scratchpad.AsSOP2()
		sp.SRC0 = 1<<32 - 9
		sp.SRC1 = 10

		alu.Run(state)

		Expect(sp.DST).To(Equal(uint64(1)))
		Expect(sp.SCC).To(Equal(uint8(1)))
	})

	It("should run S_MIN_U32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.SOP2
		state.inst.Opcode = 7

		sp := state.scratchpad.AsSOP2()
		sp.SRC0 = 1
		sp.SRC1 = 2

		alu.Run(state)

		Expect(sp.DST).To(Equal(uint64(1)))
		Expect(sp.SCC).To(Equal(uint8(1)))
	})

	It("should run S_ADD_I32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.SOP2
		state.inst.Opcode = 2

		sp := state.scratchpad.AsSOP2()
		sp.SRC0 = 1<<31 - 9
		sp.SRC1 = 10

		alu.Run(state)

		Expect(sp.DST).To(Equal(uint64(2147483649)))
		Expect(sp.SCC).To(Equal(uint8(1)))
	})

	It("should run S_ADDC_U32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.SOP2
		state.inst.Opcode = 4

		sp := state.scratchpad.AsSOP2()
		sp.SRC0 = 1<<32 - 9
		sp.SRC1 = 10
		sp.SCC = 1

		alu.Run(state)

		Expect(sp.DST).To(Equal(uint64(2)))
		Expect(sp.SCC).To(Equal(uint8(1)))
	})

	It("should run S_ASHR_I32 (Negative)", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.SOP2
		state.inst.Opcode = 34

		sp := state.Scratchpad().AsSOP2()
		sp.SRC0 = int64ToBits(-128)
		sp.SRC1 = 2

		alu.Run(state)

		Expect(sp.DST).To(Equal(uint64(int32ToBits(-32))))
		Expect(sp.SCC).To(Equal(uint8(1)))
	})

	It("should run S_ASHR_I32 (Positive)", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.SOP2
		state.inst.Opcode = 34

		sp := state.Scratchpad().AsSOP2()
		sp.SRC0 = int64ToBits(128)
		sp.SRC1 = 2

		alu.Run(state)

		Expect(sp.DST).To(Equal(uint64(int32ToBits(32))))
		Expect(sp.SCC).To(Equal(uint8(1)))
	})

	It("should run S_XOR_B32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.SOP2
		state.inst.Opcode = 18

		sp := state.Scratchpad().AsSOP2()
		sp.SRC0 = 0xf0
		sp.SRC1 = 0xff

		alu.Run(state)

		Expect(sp.DST).To(Equal(uint64(0x0f)))
		Expect(sp.SCC).To(Equal(byte(1)))
	})

})
