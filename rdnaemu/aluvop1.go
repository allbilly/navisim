package rdnaemu

import (
	"log"
)

//nolint:gocyclo,funlen
func (u *ALUImpl) runVOP1(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {

	default:
		log.Panicf("Opcode %d for VOP1 format is not implemented", inst.Opcode)
	}
}
