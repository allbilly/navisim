package rdnaemu

import "log"

//nolint:gocyclo
func (u *ALUImpl) runSOP1(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {
	case 3:
		u.runSMOVB32(state)
	case 4:
		u.runSMOVB64(state)
	case 11:
		u.runSBREVB32(state)
	case 31:
		u.runSGETPCB64(state)
	case 36:
		u.runSANDSAVEEXECB64(state)
	case 37:
		u.runSORSAVEEXECB64(state)
	case 38:
		u.runSXORSAVEEXECB64(state)
	case 39:
		u.runSANDN2SAVEEXECB64(state)
	case 40:
		u.runSORN2SAVEEXECB64(state)
	case 41:
		u.runSNANDSAVEEXECB64(state)
	case 42:
		u.runSNORSAVEEXECB64(state)
	case 43:
		u.runSNXORSAVEEXECB64(state)
	case 60:
		u.runSANDSAVEEXECB32(state)
	default:
		log.Panicf("Opcode %d for SOP1 format is not implemented", inst.Opcode)
	}
}

func (u *ALUImpl) runSMOVB32(state InstEmuState) {
	sp := state.Scratchpad().AsSOP1()
	sp.DST = sp.SRC0
}

func (u *ALUImpl) runSMOVB64(state InstEmuState) {
	sp := state.Scratchpad().AsSOP1()
	sp.DST = sp.SRC0
}

func (u *ALUImpl) runSBREVB32(state InstEmuState) {
	sp := state.Scratchpad().AsSOP1()
	dst := uint32(0)
	for i := 0; i < 32; i++ {
		bit := uint32(1 << (31 - i))
		bit = uint32(sp.SRC0) & bit
		bit = bit >> (31 - i)
		bit = bit << i
		dst = dst | bit
	}
	sp.DST = uint64(dst)
}

func (u *ALUImpl) runSGETPCB64(state InstEmuState) {
	sp := state.Scratchpad().AsSOP1()
	sp.DST = sp.PC + 4
}

func (u *ALUImpl) runSANDSAVEEXECB64(state InstEmuState) {
	sp := state.Scratchpad().AsSOP1()
	sp.DST = sp.EXEC
	sp.EXEC = sp.SRC0 & sp.EXEC
	if sp.EXEC != 0 {
		sp.SCC = 1
	} else {
		sp.SCC = 0
	}
}

func (u *ALUImpl) runSORSAVEEXECB64(state InstEmuState) {
	sp := state.Scratchpad().AsSOP1()
	sp.DST = sp.EXEC
	sp.EXEC = sp.SRC0 | sp.EXEC
	if sp.EXEC != 0 {
		sp.SCC = 1
	} else {
		sp.SCC = 0
	}
}

func (u *ALUImpl) runSXORSAVEEXECB64(state InstEmuState) {
	sp := state.Scratchpad().AsSOP1()
	sp.DST = sp.EXEC
	sp.EXEC = sp.SRC0 ^ sp.EXEC
	if sp.EXEC != 0 {
		sp.SCC = 1
	} else {
		sp.SCC = 0
	}
}

func (u *ALUImpl) runSANDN2SAVEEXECB64(state InstEmuState) {
	sp := state.Scratchpad().AsSOP1()
	sp.DST = sp.EXEC
	sp.EXEC = sp.SRC0 & (^sp.EXEC)
	if sp.EXEC != 0 {
		sp.SCC = 1
	} else {
		sp.SCC = 0
	}
}

func (u *ALUImpl) runSORN2SAVEEXECB64(state InstEmuState) {
	sp := state.Scratchpad().AsSOP1()
	sp.DST = sp.EXEC
	sp.EXEC = sp.SRC0 | (^sp.EXEC)
	if sp.EXEC != 0 {
		sp.SCC = 1
	} else {
		sp.SCC = 0
	}
}

func (u *ALUImpl) runSNANDSAVEEXECB64(state InstEmuState) {
	sp := state.Scratchpad().AsSOP1()
	sp.DST = sp.EXEC
	sp.EXEC = ^(sp.SRC0 & sp.EXEC)
	if sp.EXEC != 0 {
		sp.SCC = 1
	} else {
		sp.SCC = 0
	}
}

func (u *ALUImpl) runSNORSAVEEXECB64(state InstEmuState) {
	sp := state.Scratchpad().AsSOP1()
	sp.DST = sp.EXEC
	sp.EXEC = ^(sp.SRC0 | sp.EXEC)
	if sp.EXEC != 0 {
		sp.SCC = 1
	} else {
		sp.SCC = 0
	}
}

func (u *ALUImpl) runSNXORSAVEEXECB64(state InstEmuState) {
	sp := state.Scratchpad().AsSOP1()
	sp.DST = sp.EXEC
	sp.EXEC = ^(sp.SRC0 ^ sp.EXEC)
	if sp.EXEC != 0 {
		sp.SCC = 1
	} else {
		sp.SCC = 0
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
