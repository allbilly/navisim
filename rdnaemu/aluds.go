package rdnaemu

import (
	"log"
)

func (u *ALUImpl) runDS(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {
	case 223:
		u.DSWRITEB128(state)
	default:
		log.Panicf("Opcode %d for DS format is not implemented", inst.Opcode)
	}
}

func (u *ALUImpl) DSWRITEB128(state InstEmuState) {

}
