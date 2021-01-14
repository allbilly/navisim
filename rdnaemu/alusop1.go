package rdnaemu

import "log"

//nolint:gocyclo
func (u *ALUImpl) runSOP1(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {

	default:
		log.Panicf("Opcode %d for SOP1 format is not implemented", inst.Opcode)
	}
}
