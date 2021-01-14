package rdnaemu

import (
	"log"
)

//nolint:gocyclo,funlen
func (u *ALUImpl) runVOP2(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {

	default:
		log.Panicf("Opcode %d for VOP2 format (%s) is not implemented",
			inst.Opcode, inst.String(nil))
	}
}
