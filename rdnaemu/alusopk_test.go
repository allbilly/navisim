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

	It("should run s_movk_i32", func() {
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.SOPK
		state.inst.Opcode = 0

		sp := state.Scratchpad().AsSOPK()
		sp.IMM = uint64(int16ToBits(-12))

		alu.Run(state)

		Expect(asInt16(uint16(sp.DST))).To(Equal(int16(-12)))
	})
})
