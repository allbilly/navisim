package rdnaemu

import (
	"log"
)

//nolint:gocyclo,funlen
func (u *ALUImpl) runSOPC(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {
	case 6:
		u.runSCMPEQU32(state)

	default:
		log.Panicf("Opcode %d for SOPC format is not implemented", inst.Opcode)
	}
}

func (u *ALUImpl) runSCMPEQU32(state InstEmuState) {
	sp := state.Scratchpad().AsSOPC()
	if uint32(sp.SRC0) == uint32(sp.SRC1) {
		sp.SCC = 1
	} else {
		sp.SCC = 0
	}
}
