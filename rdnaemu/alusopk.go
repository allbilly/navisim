package rdnaemu

import (
	"log"
)

func (u *ALUImpl) runSOPK(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {
	case 23:
		//SWAITCNTVSCNT
	default:
		log.Panicf("Opcode %d for SOPK format is not implemented", inst.Opcode)
	}
}
