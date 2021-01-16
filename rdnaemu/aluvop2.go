package rdnaemu

import (
	"log"
	"math"
)

//nolint:gocyclo,funlen
func (u *ALUImpl) runVOP2(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {
	case 1:
		u.runVCNDMASKB32(state)
	case 15:
		u.runVMINF32(state)
	case 16:
		u.runVMAXF32(state)
	case 37:
		u.runVADDNCU32(state)
	case 38:
		u.runVSUBNCU32(state)
	case 39:
		u.runVSUBREVNCU32(state)
	case 40:
		u.runVADDCOCIU32(state)
	case 43:
		u.runVFMACF32(state)

	default:
		log.Panicf("Opcode %d for VOP2 format (%s) is not implemented",
			inst.Opcode, inst.String(nil))
	}
}

func (u *ALUImpl) runVCNDMASKB32(state InstEmuState) {
	sp := state.Scratchpad().AsVOP2()
	inst := state.Inst()
	if inst.IsSdwa == false {
		var i uint
		for i = 0; i < 64; i++ {
			if !laneMasked(sp.EXEC, i) {
				continue
			}

			if (sp.VCC & (1 << i)) > 0 {
				sp.DST[i] = sp.SRC1[i]
			} else {
				sp.DST[i] = sp.SRC0[i]
			}
		}
	} else {
		log.Panicf("SDWA for VOP2 instruction opcode %d not implemented \n", inst.Opcode)
	}

}

func (u *ALUImpl) runVADDNCU32(state InstEmuState) {
	sp := state.Scratchpad().AsVOP2()
	inst := state.Inst()
	var i uint
	if !inst.IsSdwa {
		for i = 0; i < 64; i++ {
			if !laneMasked(sp.EXEC, i) {
				continue
			}
			src0 := uint32(sp.SRC0[i])
			src1 := uint32(sp.SRC1[i])
			dst := src0 + src1
			sp.DST[i] = uint64(dst)
		}
	} else {
		for i = 0; i < 64; i++ {
			if !laneMasked(sp.EXEC, i) {
				continue
			}
			src0 := u.sdwaSrcSelect(uint32(sp.SRC0[i]), inst.Src0Sel)
			src1 := u.sdwaSrcSelect(uint32(sp.SRC1[i]), inst.Src1Sel)
			dst := src0 + src1
			dst = u.sdwaDstSelect(uint32(sp.DST[i]), dst,
				inst.DstSel, inst.DstUnused)
			sp.DST[i] = uint64(dst)
		}
	}
}
func (u *ALUImpl) runVSUBNCU32(state InstEmuState) {
	sp := state.Scratchpad().AsVOP2()
	inst := state.Inst()
	var i uint
	if !inst.IsSdwa {
		for i = 0; i < 64; i++ {
			if !laneMasked(sp.EXEC, i) {
				continue
			}
			src0 := uint32(sp.SRC0[i])
			src1 := uint32(sp.SRC1[i])
			dst := src0 - src1
			sp.DST[i] = uint64(dst)
		}
	} else {
		for i = 0; i < 64; i++ {
			if !laneMasked(sp.EXEC, i) {
				continue
			}
			src0 := u.sdwaSrcSelect(uint32(sp.SRC0[i]), inst.Src0Sel)
			src1 := u.sdwaSrcSelect(uint32(sp.SRC1[i]), inst.Src1Sel)
			dst := src0 - src1
			dst = u.sdwaDstSelect(uint32(sp.DST[i]), dst,
				inst.DstSel, inst.DstUnused)
			sp.DST[i] = uint64(dst)
		}
	}
}

func (u *ALUImpl) runVSUBREVNCU32(state InstEmuState) {
	sp := state.Scratchpad().AsVOP2()
	inst := state.Inst()
	var i uint
	if !inst.IsSdwa {
		for i = 0; i < 64; i++ {
			if !laneMasked(sp.EXEC, i) {
				continue
			}
			src0 := uint32(sp.SRC0[i])
			src1 := uint32(sp.SRC1[i])
			dst := src1 - src0
			sp.DST[i] = uint64(dst)
		}
	} else {
		for i = 0; i < 64; i++ {
			if !laneMasked(sp.EXEC, i) {
				continue
			}
			src0 := u.sdwaSrcSelect(uint32(sp.SRC0[i]), inst.Src0Sel)
			src1 := u.sdwaSrcSelect(uint32(sp.SRC1[i]), inst.Src1Sel)
			dst := src1 - src0
			dst = u.sdwaDstSelect(uint32(sp.DST[i]), dst,
				inst.DstSel, inst.DstUnused)
			sp.DST[i] = uint64(dst)
		}
	}
}
func (u *ALUImpl) runVADDCOCIU32(state InstEmuState) {
	sp := state.Scratchpad().AsVOP2()
	inst := state.Inst()
	var i uint
	if !inst.IsSdwa {
		for i = 0; i < 64; i++ {
			if !laneMasked(sp.EXEC, i) {
				continue
			}
			src0 := uint32(sp.SRC0[i])
			src1 := uint32(sp.SRC1[i])
			dst := src0 + src1 + uint32(sp.VCC)
			sp.DST[i] = uint64(dst)
			if src0 > math.MaxUint32-src1 {
				sp.VCC = 1
			} else {
				sp.VCC = 0
			}
		}
	} else {
		for i = 0; i < 64; i++ {
			if !laneMasked(sp.EXEC, i) {
				continue
			}
			src0 := u.sdwaSrcSelect(uint32(sp.SRC0[i]), inst.Src0Sel)
			src1 := u.sdwaSrcSelect(uint32(sp.SRC1[i]), inst.Src1Sel)
			dst := src0 + src1 + uint32(sp.VCC)
			dst = u.sdwaDstSelect(uint32(sp.DST[i]), dst,
				inst.DstSel, inst.DstUnused)
			sp.DST[i] = uint64(dst)
			if src0 > math.MaxUint32-src1 {
				sp.VCC = 1
			} else {
				sp.VCC = 0
			}
		}
	}
}
func (u *ALUImpl) runVMINF32(state InstEmuState) {
	sp := state.Scratchpad().AsVOP2()
	inst := state.Inst()
	if !inst.IsSdwa {
		var i uint
		for i = 0; i < 64; i++ {
			if !laneMasked(sp.EXEC, i) {
				continue
			}

			src0 := math.Float32frombits(uint32(sp.SRC0[i]))
			src1 := math.Float32frombits(uint32(sp.SRC1[i]))
			dst := src0
			if src1 < src0 {
				dst = src1
			}

			sp.DST[i] = uint64(math.Float32bits(dst))
		}
	} else {
		log.Panicf("SDWA for VOP2 instruction opcode %d not implemented \n", inst.Opcode)
	}
}

func (u *ALUImpl) runVMAXF32(state InstEmuState) {
	sp := state.Scratchpad().AsVOP2()
	inst := state.Inst()
	if !inst.IsSdwa {
		var i uint
		for i = 0; i < 64; i++ {
			if !laneMasked(sp.EXEC, i) {
				continue
			}

			src0 := math.Float32frombits(uint32(sp.SRC0[i]))
			src1 := math.Float32frombits(uint32(sp.SRC1[i]))
			dst := src0
			if src1 > src0 {
				dst = src1
			}

			sp.DST[i] = uint64(math.Float32bits(dst))
		}
	} else {
		log.Panicf("SDWA for VOP2 instruction opcode %d not implemented \n", inst.Opcode)
	}
}

func (u *ALUImpl) runVFMACF32(state InstEmuState) {
	sp := state.Scratchpad().AsVOP2()
	inst := state.Inst()
	var dst float32
	var src0 float32
	var src1 float32

	var i uint
	if inst.IsSdwa == false {
		for i = 0; i < 64; i++ {
			if !laneMasked(sp.EXEC, i) {
				continue
			}

			dst = asFloat32(uint32(sp.DST[i]))
			src0 = asFloat32(uint32(sp.SRC0[i]))
			src1 = asFloat32(uint32(sp.SRC1[i]))
			dst += src0 * src1
			sp.DST[i] = uint64(float32ToBits(dst))
		}
	} else {
		log.Panicf("SDWA for VOP2 instruction opcode  %d not implemented \n", inst.Opcode)
	}
}
