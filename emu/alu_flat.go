package emu

import (
	"log"
)

//nolint:gocyclo
//nolint:funlen
func (u *ALUImpl) runFlat(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {

	default:
		log.Panicf("Opcode %d for FLAT format is not implemented", inst.Opcode)
	}
}
