package rdnaemu

import (
	"log"
)

//nolint:gocyclo,funlen
func (u *ALUImpl) runSOPC(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {
	default:
		log.Panicf("Opcode %d for SOPC format is not implemented", inst.Opcode)
	}
}
