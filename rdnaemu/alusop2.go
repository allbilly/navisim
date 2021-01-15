package rdnaemu

import (
	"log"
	"math"
)

//nolint:gocyclo,funlen
func (u *ALUImpl) runSOP2(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {
	case 0:
		u.runSADDU32(state)
	case 2:
		u.runSADDI32(state)
	case 14:
		u.runSANDB32(state)
	case 15:
		u.runSANDB64(state)
	case 38:
		u.runSMULI32(state)
	default:
		log.Panicf("Opcode %d for SOP2 format is not implemented", inst.Opcode)
	}
}

func (u *ALUImpl) runSADDU32(state InstEmuState) {
	sp := state.Scratchpad().AsSOP2()

	src0 := uint32(sp.SRC0)
	src1 := uint32(sp.SRC1)
	dst := src0 + src1

	if src0 > math.MaxUint32-src1 {
		sp.SCC = 1
	} else {
		sp.SCC = 0
	}

	sp.DST = uint64(dst)

}

func (u *ALUImpl) runSADDI32(state InstEmuState) {
	sp := state.Scratchpad().AsSOP2()

	src0 := asInt32(uint32(sp.SRC0))
	src1 := asInt32(uint32(sp.SRC1))
	dst := src0 + src1

	if src0 > 0 && src1 > 0 && dst < 0 {
		sp.SCC = 1
	} else if src0 < 0 && src1 < 0 && dst < 0 {
		sp.SCC = 1
	} else {
		sp.SCC = 0
	}

	sp.DST = uint64(int32ToBits(dst))

}

func (u *ALUImpl) runSANDB32(state InstEmuState) {
	sp := state.Scratchpad().AsSOP2()

	sp.DST = sp.SRC0 & sp.SRC1
	if sp.DST != 0 {
		sp.SCC = 1
	} else {
		sp.SCC = 0
	}
}
func (u *ALUImpl) runSANDB64(state InstEmuState) {
	sp := state.Scratchpad().AsSOP2()

	sp.DST = sp.SRC0 & sp.SRC1
	if sp.DST != 0 {
		sp.SCC = 1
	} else {
		sp.SCC = 0
	}
}
func (u *ALUImpl) runSMULI32(state InstEmuState) {
	sp := state.Scratchpad().AsSOP2()

	src0 := asInt32(uint32(sp.SRC0))
	src1 := asInt32(uint32(sp.SRC1))
	dst := src0 * src1

	sp.DST = uint64(int32ToBits(dst))

	if src0 != 0 && dst/src0 != src1 {
		sp.SCC = 1
	}
}
