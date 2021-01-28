package rdnaemu

import (
	"log"
)

//nolint:gocyclo,funlen
func (u *ALUImpl) runVOPC(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {
	case 129: // v_cmp_gt_i32_e32
		u.runVCMPLTI32(state)
	case 132: // v_cmp_gt_i32_e32
		u.runVCMPGTI32(state)
	case 194:
		u.runVCMPEQU32(state)
	case 195:
		u.runVCMPLEU32(state)
	case 196:
		u.runVCMPGTU32(state)
	case 197:
		u.runVCMPNEU32(state)
	case 198:
		u.runVCMPGEU32(state)

	default:
		log.Panicf("Opcode %d for VOPC format is not implemented", inst.Opcode)
	}
}

func (u *ALUImpl) runVCMPLTI32(state InstEmuState) {
	sp := state.Scratchpad().AsVOPC()
	sp.VCC = 0
	var i uint
	for i = 0; i < 64; i++ {
		if !laneMasked(sp.EXEC, i) {
			continue
		}

		if sp.SRC0[i] < sp.SRC1[i] {
			sp.VCC = sp.VCC | (1 << i)
		}
	}
}

func (u *ALUImpl) runVCMPGTI32(state InstEmuState) {
	sp := state.Scratchpad().AsVOPC()
	sp.VCC = 0
	var i uint
	for i = 0; i < 64; i++ {
		if !laneMasked(sp.EXEC, i) {
			continue
		}

		src0 := asInt32(uint32(sp.SRC0[i]))
		src1 := asInt32(uint32(sp.SRC1[i]))
		if src0 > src1 {
			sp.VCC = sp.VCC | (1 << i)
		}
	}
}

func (u *ALUImpl) runVCMPEQU32(state InstEmuState) {
	sp := state.Scratchpad().AsVOPC()
	sp.VCC = 0
	var i uint
	for i = 0; i < 64; i++ {
		if !laneMasked(sp.EXEC, i) {
			continue
		}

		if uint32(sp.SRC0[i]) == uint32(sp.SRC1[i]) {
			sp.VCC = sp.VCC | (1 << i)
		}
	}
}

func (u *ALUImpl) runVCMPLEU32(state InstEmuState) {
	sp := state.Scratchpad().AsVOPC()
	sp.VCC = 0
	var i uint
	for i = 0; i < 64; i++ {
		if !laneMasked(sp.EXEC, i) {
			continue
		}

		src0 := uint32(sp.SRC0[i])
		src1 := uint32(sp.SRC1[i])
		if src0 <= src1 {
			sp.VCC = sp.VCC | (1 << i)
		}
	}
}

func (u *ALUImpl) runVCMPGTU32(state InstEmuState) {
	sp := state.Scratchpad().AsVOPC()
	sp.VCC = 0
	var i uint
	for i = 0; i < 64; i++ {
		if !laneMasked(sp.EXEC, i) {
			continue
		}

		src0 := uint32(sp.SRC0[i])
		src1 := uint32(sp.SRC1[i])
		if src0 > src1 {
			sp.VCC = sp.VCC | (1 << i)
		}
	}
}

func (u *ALUImpl) runVCMPGEU32(state InstEmuState) {
	sp := state.Scratchpad().AsVOPC()
	sp.VCC = 0
	var i uint
	for i = 0; i < 64; i++ {
		if !laneMasked(sp.EXEC, i) {
			continue
		}

		src0 := uint32(sp.SRC0[i])
		src1 := uint32(sp.SRC1[i])
		if src0 >= src1 {
			sp.VCC = sp.VCC | (1 << i)
		}
	}
}

func (u *ALUImpl) runVCMPNEU32(state InstEmuState) {
	sp := state.Scratchpad().AsVOPC()
	sp.VCC = 0
	var i uint
	for i = 0; i < 64; i++ {
		if laneMasked(sp.EXEC, i) {
			if sp.SRC0[i] != sp.SRC1[i] {
				sp.VCC = sp.VCC | (1 << i)
			}
		}
	}
}
