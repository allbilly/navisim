package emu

import (
	"log"

	"gitlab.com/akita/navisim/insts"
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
	// fmt.Printf("%s\n", inst.String(nil))

	switch inst.FormatType {
	case insts.SOP1:
		u.runSOP1(state)
	case insts.SOP2:
		u.runSOP2(state)
	case insts.SOPC:
		u.runSOPC(state)
	case insts.SMEM:
		u.runSMEM(state)
	case insts.VOP1:
		u.runVOP1(state)
	case insts.VOP2:
		u.runVOP2(state)
	case insts.VOP3a:
		u.runVOP3A(state)
	case insts.VOP3b:
		u.runVOP3B(state)
	case insts.VOPC:
		u.runVOPC(state)
	case insts.FLAT:
		u.runFlat(state)
	case insts.SOPP:
		u.runSOPP(state)
	case insts.SOPK:
		u.runSOPK(state)
	case insts.DS:
		u.runDS(state)
	default:
		log.Panicf("Inst format %s is not supported", inst.Format.FormatName)
	}
}

func (u *ALUImpl) runSMEM(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {

	default:
		log.Panicf("Opcode %d for SMEM format is not implemented", inst.Opcode)
	}
}

//nolint:gocyclo
func (u *ALUImpl) runSOPP(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {

	default:
		log.Panicf("Opcode %d for SOPP format is not implemented", inst.Opcode)
	}
}

func laneMasked(Exec uint64, laneID uint) bool {
	return Exec&(1<<laneID) > 0
}
