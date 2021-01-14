package emu

import (
	"log"
	"math"
	"strings"
)

//nolint:gocyclo,funlen
func (u *ALUImpl) runVOP3A(state InstEmuState) {
	inst := state.Inst()

	u.vop3aPreprocess(state)

	switch inst.Opcode {

	default:
		log.Panicf("Opcode %d for VOP3a format is not implemented", inst.Opcode)
	}
	u.vop3aPostprocess(state)
}

func (u *ALUImpl) vop3aPreprocess(state InstEmuState) {
	inst := state.Inst()

	if inst.Abs != 0 {
		u.vop3aPreProcessAbs(state)
	}

	if inst.Neg != 0 {
		u.vop3aPreProcessNeg(state)
	}
}

func (u *ALUImpl) vop3aPreProcessAbs(state InstEmuState) {
	inst := state.Inst()
	sp := state.Scratchpad().AsVOP3A()

	if strings.Contains(inst.InstName, "F32") ||
		strings.Contains(inst.InstName, "f32") {
		if inst.Abs&0x1 != 0 {
			for i := 0; i < 64; i++ {
				src0 := math.Float32frombits(uint32(sp.SRC0[i]))
				src0 = float32(math.Abs(float64(src0)))
				sp.SRC0[i] = uint64(math.Float32bits(src0))
			}
		}

		if inst.Abs&0x2 != 0 {
			for i := 0; i < 64; i++ {
				src1 := math.Float32frombits(uint32(sp.SRC1[i]))
				src1 = float32(math.Abs(float64(src1)))
				sp.SRC1[i] = uint64(math.Float32bits(src1))
			}
		}

		if inst.Abs&0x4 != 0 {
			for i := 0; i < 64; i++ {
				src2 := math.Float32frombits(uint32(sp.SRC2[i]))
				src2 = float32(math.Abs(float64(src2)))
				sp.SRC2[i] = uint64(math.Float32bits(src2))
			}
		}
	} else {
		log.Printf("Absolute operation for %s is not implemented.", inst.InstName)
	}
}

func (u *ALUImpl) vop3aPreProcessNeg(state InstEmuState) {
	inst := state.Inst()

	if strings.Contains(inst.InstName, "F64") ||
		strings.Contains(inst.InstName, "f64") {
		u.vop3aPreProcessF64Neg(state)
	} else if strings.Contains(inst.InstName, "F32") ||
		strings.Contains(inst.InstName, "f32") {
		u.vop3aPreProcessF32Neg(state)
	} else if strings.Contains(inst.InstName, "B32") ||
		strings.Contains(inst.InstName, "b32") {
		u.vop3aPreProcessB32Neg(state)
	} else {
		log.Printf("Negative operation for %s is not implemented.", inst.InstName)
	}
}

func (u *ALUImpl) vop3aPreProcessF64Neg(state InstEmuState) {
	inst := state.Inst()
	sp := state.Scratchpad().AsVOP3A()

	if inst.Neg&0x1 != 0 {
		for i := 0; i < 64; i++ {
			src0 := math.Float64frombits(sp.SRC0[i])
			src0 = src0 * (-1.0)
			sp.SRC0[i] = math.Float64bits(src0)
		}
	}

	if inst.Neg&0x2 != 0 {
		for i := 0; i < 64; i++ {
			src1 := math.Float64frombits(sp.SRC1[i])
			src1 = src1 * (-1.0)
			sp.SRC1[i] = math.Float64bits(src1)
		}
	}

	if inst.Neg&0x4 != 0 {
		for i := 0; i < 64; i++ {
			src2 := math.Float64frombits(sp.SRC2[i])
			src2 = src2 * (-1.0)
			sp.SRC2[i] = math.Float64bits(src2)
		}
	}
}

func (u *ALUImpl) vop3aPreProcessF32Neg(state InstEmuState) {
	inst := state.Inst()
	sp := state.Scratchpad().AsVOP3A()
	if inst.Neg&0x1 != 0 {
		for i := 0; i < 64; i++ {
			src0 := math.Float32frombits(uint32(sp.SRC0[i]))
			src0 = src0 * (-1.0)
			sp.SRC0[i] = uint64(math.Float32bits(src0))
		}
	}

	if inst.Neg&0x2 != 0 {
		for i := 0; i < 64; i++ {
			src1 := math.Float32frombits(uint32(sp.SRC1[i]))
			src1 = src1 * (-1.0)
			sp.SRC1[i] = uint64(math.Float32bits(src1))
		}
	}

	if inst.Neg&0x4 != 0 {
		for i := 0; i < 64; i++ {
			src2 := math.Float32frombits(uint32(sp.SRC2[i]))
			src2 = src2 * (-1.0)
			sp.SRC2[i] = uint64(math.Float32bits(src2))
		}
	}
}

func (u *ALUImpl) vop3aPreProcessB32Neg(state InstEmuState) {
	inst := state.Inst()
	sp := state.Scratchpad().AsVOP3A()
	if inst.Neg&0x1 != 0 {
		for i := 0; i < 64; i++ {
			src0 := asInt32(uint32(sp.SRC0[i]))
			src0 = src0 * (-1.0)
			sp.SRC0[i] = uint64(int32ToBits(src0))
		}
	}

	if inst.Neg&0x2 != 0 {
		for i := 0; i < 64; i++ {
			src1 := asInt32(uint32(sp.SRC1[i]))
			src1 = src1 * (-1.0)
			sp.SRC1[i] = uint64(int32ToBits(src1))
		}
	}

	if inst.Neg&0x4 != 0 {
		for i := 0; i < 64; i++ {
			src2 := asInt32(uint32(sp.SRC2[i]))
			src2 = src2 * (-1.0)
			sp.SRC2[i] = uint64(int32ToBits(src2))
		}
	}
}

func (u *ALUImpl) vop3aPostprocess(state InstEmuState) {
	inst := state.Inst()

	if inst.Omod != 0 {
		log.Panic("Output modifiers are not supported.")
	}
}
