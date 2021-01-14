package emu

import (
	"log"
)

//nolint:gocyclo,funlen
func (u *ALUImpl) runVOPC(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {

	default:
		log.Panicf("Opcode 0x%02X for VOPC format is not implemented", inst.Opcode)
	}
}
