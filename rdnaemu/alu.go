package rdnaemu

import (
	"log"

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

	default:
		log.Panicf("Opcode %d for SMEM format is not implemented", inst.Opcode)
	}
}
