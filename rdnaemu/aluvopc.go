package rdnaemu

import (
	"log"
)

//nolint:gocyclo,funlen
func (u *ALUImpl) runVOPC(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {
	case 132: // v_cmp_gt_i32_e32
		u.runVCmpGtI32(state)

	default:
		log.Panicf("Opcode %d for VOPC format is not implemented", inst.Opcode)
	}
}
func (u *ALUImpl) runVCmpGtI32(state InstEmuState) {
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
