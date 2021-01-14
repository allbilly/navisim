package rdnaemu

import "log"

//nolint:gocyclo
func (u *ALUImpl) runSOP1(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {
	case 3:
		u.runSMOVB32(state)
	case 60:
		u.runSANDSAVEEXECB32(state)

	default:
		log.Panicf("Opcode %d for SOP1 format is not implemented", inst.Opcode)
	}
}
func (u *ALUImpl) runSANDSAVEEXECB32(state InstEmuState) {
	sp := state.Scratchpad().AsSOP1()
	sp.DST = sp.EXEC
	sp.EXEC = (sp.SRC0 & sp.EXEC)
	if sp.EXEC != 0 {
		sp.SCC = 1
	} else {
		sp.SCC = 0
	}
}

func (u *ALUImpl) runSMOVB32(state InstEmuState) {
	sp := state.Scratchpad().AsSOP1()
	sp.DST = sp.SRC0
}
