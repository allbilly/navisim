package rdnaemu

import (
	"log"
)

//nolint:gocyclo,funlen
func (u *ALUImpl) runVOP1(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {
	case 1:
		u.runVMOVB32(state)

	default:
		log.Panicf("Opcode %d for VOP1 format is not implemented", inst.Opcode)
	}
}
func (u *ALUImpl) runVMOVB32(state InstEmuState) {
	sp := state.Scratchpad().AsVOP1()
	var i uint
	for i = 0; i < 64; i++ {
		if !laneMasked(sp.EXEC, i) {
			continue
		}

		sp.DST[i] = sp.SRC0[i]
	}
}
