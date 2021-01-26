package rdnaemu

import (
	"log"
	"math"
)

//nolint:gocyclo,funlen
func (u *ALUImpl) runVOP1(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {
	case 1:
		u.runVMOVB32(state)
	case 6:
		u.runVCVTF32U32(state)
	case 43:
		u.runVRCPIFLAGF32(state)

	default:
		log.Panicf("Opcode %d for VOP1 format is not implemented", inst.Opcode)
	}
}
func (u *ALUImpl) runVMOVB32(state InstEmuState) {
	sp := state.Scratchpad().AsVOP1()
	var i uint
	for i = 0; i < 64; i++ {
		if !laneMasked(sp.EXEC, i) {
			continue
		}

		sp.DST[i] = sp.SRC0[i]
	}
}
func (u *ALUImpl) runVCVTF32U32(state InstEmuState) {
	sp := state.Scratchpad().AsVOP1()

	var i uint
	for i = 0; i < 64; i++ {
		if !laneMasked(sp.EXEC, i) {
			continue
		}
		sp.DST[i] = uint64(math.Float32bits(float32(uint32(sp.SRC0[i]))))
	}
}
func (u *ALUImpl) runVRCPIFLAGF32(state InstEmuState) {
	sp := state.Scratchpad().AsVOP1()

	var i uint
	for i = 0; i < 64; i++ {
		if !laneMasked(sp.EXEC, i) {
			continue
		}

		src := math.Float32frombits(uint32(sp.SRC0[i]))
		dst := 1 / src
		sp.DST[i] = uint64(math.Float32bits(dst))
	}
}
