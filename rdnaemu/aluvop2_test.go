package rdnaemu

import (
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("ALU", func() {

	var (
		// alu   *ALUImpl
		state *mockInstState
	)

	BeforeEach(func() {
		// alu = NewALU(nil)

		state = new(mockInstState)
		state.scratchpad = make([]byte, 4096)
	})

	// It("should run V_LSHRREV_B32", func() {
	// 	state.inst = rdnainsts.NewInst()
	// 	state.inst.FormatType = rdnainsts.VOP2
	// 	state.inst.Opcode = 26

	// 	sp := state.Scratchpad().AsVOP2()
	// 	sp.SRC0[0] = 0x64
	// 	sp.SRC1[0] = 0x02
	// 	sp.EXEC = 0x1

	// 	alu.Run(state)

	// 	Expect(uint32(sp.DST[0])).To(Equal(uint32(0x20)))
	// })

})
