package rdnaemu

import (
	"log"
)

//nolint:gocyclo,funlen
func (u *ALUImpl) runSOP2(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {

	default:
		log.Panicf("Opcode %d for SOP2 format is not implemented", inst.Opcode)
	}
}
