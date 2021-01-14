package emu

import (
	"log"
)

func (u *ALUImpl) runSOPK(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {

	default:
		log.Panicf("Opcode %d for SOPK format is not implemented", inst.Opcode)
	}
}
