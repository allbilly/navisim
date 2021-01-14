package rdnaemu

import (
	"log"
)

func (u *ALUImpl) runDS(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {

	default:
		log.Panicf("Opcode %d for DS format is not implemented", inst.Opcode)
	}
}
