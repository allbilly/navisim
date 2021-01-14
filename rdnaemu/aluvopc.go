package rdnaemu

import (
	"log"
)

//nolint:gocyclo,funlen
func (u *ALUImpl) runVOPC(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {

	default:
		log.Panicf("Opcode %d for VOPC format is not implemented", inst.Opcode)
	}
}
