package rdnaemu

import (
	"bytes"
	"fmt"
	"log"

	"encoding/binary"

	"gitlab.com/akita/navisim/rdnainsts"
)

//ALU does its jobs
type ALU interface {
	Run(state InstEmuState)

	SetLDS(lds []byte)
	LDS() []byte
}

// ALUImpl is where the instructions get executed.
type ALUImpl struct {
	storageAccessor *storageAccessor
	lds             []byte
}

// NewALU creates a new ALU with a storage as a dependency.
func NewALU(storageAccessor *storageAccessor) *ALUImpl {
	alu := new(ALUImpl)
	alu.storageAccessor = storageAccessor
	return alu
}

// SetLDS assigns the LDS storage to be used in the following instructions.
func (u *ALUImpl) SetLDS(lds []byte) {
	u.lds = lds
}

//LDS returns lds
func (u *ALUImpl) LDS() []byte {
	return u.lds
}

// Run executes the instruction in the scatchpad of the InstEmuState
//nolint:gocyclo
func (u *ALUImpl) Run(state InstEmuState) {
	inst := state.Inst()
	//fmt.Printf("%s\n", inst.String(nil))

	switch inst.FormatType {
	case rdnainsts.SOP1:
		u.runSOP1(state)
	case rdnainsts.SOP2:
		u.runSOP2(state)
	case rdnainsts.SOPC:
		u.runSOPC(state)
	case rdnainsts.SMEM:
		u.runSMEM(state)
	case rdnainsts.VOP1:
		u.runVOP1(state)
	case rdnainsts.VOP2:
		u.runVOP2(state)
	case rdnainsts.VOP3a:
		u.runVOP3A(state)
	case rdnainsts.VOP3b:
		u.runVOP3B(state)
	case rdnainsts.VOPC:
		u.runVOPC(state)
	case rdnainsts.FLAT:
		u.runFlat(state)
	case rdnainsts.SOPP:
		u.runSOPP(state)
	case rdnainsts.SOPK:
		u.runSOPK(state)
	case rdnainsts.DS:
		u.runDS(state)
	default:
		log.Panicf("Inst format %s is not supported", inst.Format.FormatName)
	}
}

func (u *ALUImpl) runSMEM(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {
	case 0:
		u.runSLOADDWORD(state)
	case 1:
		u.runSLOADDWORDX2(state)
	case 2:
		u.runSLOADDWORDX4(state)
	case 3:
		u.runSLOADDWORDX8(state)
	default:
		log.Panicf("Opcode %d for SMEM format is not implemented", inst.Opcode)
	}
}

func (u *ALUImpl) runSLOADDWORD(state InstEmuState) {
	sp := state.Scratchpad().AsSMEM()
	pid := state.PID()

	buf := u.storageAccessor.Read(pid, sp.Base+sp.Offset, 4)

	sp.DST[0] = rdnainsts.BytesToUint32(buf)
}

func (u *ALUImpl) runSLOADDWORDX2(state InstEmuState) {
	sp := state.Scratchpad().AsSMEM()
	spRaw := state.Scratchpad()
	pid := state.PID()

	buf := u.storageAccessor.Read(pid, sp.Base+sp.Offset, 8)
	copy(spRaw[32:40], buf)
}

func (u *ALUImpl) runSLOADDWORDX4(state InstEmuState) {
	sp := state.Scratchpad().AsSMEM()
	spRaw := state.Scratchpad()
	pid := state.PID()

	buf := u.storageAccessor.Read(pid, sp.Base+sp.Offset, 16)
	copy(spRaw[32:48], buf)
}

func (u *ALUImpl) runSLOADDWORDX8(state InstEmuState) {
	sp := state.Scratchpad().AsSMEM()
	spRaw := state.Scratchpad()
	pid := state.PID()

	buf := u.storageAccessor.Read(pid, sp.Base+sp.Offset, 32)
	copy(spRaw[32:64], buf)
}

func (u *ALUImpl) sdwaSrcSelect(src uint32, sel rdnainsts.SDWASelect) uint32 {
	switch sel {
	case rdnainsts.SDWASelectByte0:
		return src & 0x000000ff
	case rdnainsts.SDWASelectByte1:
		return (src & 0x0000ff00) >> 8
	case rdnainsts.SDWASelectByte2:
		return (src & 0x00ff0000) >> 16
	case rdnainsts.SDWASelectByte3:
		return (src & 0xff000000) >> 24
	case rdnainsts.SDWASelectWord0:
		return src & 0x0000ffff
	case rdnainsts.SDWASelectWord1:
		return (src & 0xffff0000) >> 16
	case rdnainsts.SDWASelectDWord:
		return src
	}
	return src
}

func (u *ALUImpl) sdwaDstSelect(
	dstOld uint32,
	dstNew uint32,
	sel rdnainsts.SDWASelect,
	unused rdnainsts.SDWAUnused,
) uint32 {
	value := dstNew
	switch sel {
	case rdnainsts.SDWASelectByte0:
		value = value & 0x000000ff
	case rdnainsts.SDWASelectByte1:
		value = (value << 8) & 0x0000ff00
	case rdnainsts.SDWASelectByte2:
		value = (value << 16) & 0x00ff0000
	case rdnainsts.SDWASelectByte3:
		value = (value << 24) & 0xff000000
	case rdnainsts.SDWASelectWord0:
		value = value & 0x0000ffff
	case rdnainsts.SDWASelectWord1:
		value = (value << 16) & 0xffff0000
	}

	return value
}

//nolint:unused
func (u *ALUImpl) dumpScratchpadAsSop2(
	state InstEmuState,
	byteCount int,
) string {
	scratchpad := state.Scratchpad()
	layout := new(SOP2Layout)

	err := binary.Read(bytes.NewBuffer(scratchpad), binary.LittleEndian, layout)
	if err != nil {
		panic(err)
	}

	output := fmt.Sprintf(
		`
			SRC0: 0x%[1]x(%[1]d),
			SRC1: 0x%[2]x(%[2]d),
			SCC: 0x%[3]x(%[3]d),
			DST: 0x%[4]x(%[4]d)\n",
		`,
		layout.SRC0, layout.SRC1, layout.SCC, layout.DST)

	return output
}
