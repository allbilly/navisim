package rdnaemu

import (
	"log"
)

func (u *ALUImpl) runDS(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {
	case 13:
		u.runDSWRITEB32(state)
	default:
		log.Panicf("Opcode %d for DS format is not implemented", inst.Opcode)
	}
}

func (u *ALUImpl) runDSWRITEB32(state InstEmuState) {
	inst := state.Inst()
	sp := state.Scratchpad()
	layout := sp.AsDS()
	lds := u.LDS()

	i := uint(0)
	for i = 0; i < 64; i++ {
		if !laneMasked(layout.EXEC, i) {
			continue
		}

		addr0 := layout.ADDR[i] + inst.Offset0
		data0offset := uint(8 + 64*4)

		copy(lds[addr0:addr0+4], sp[data0offset+i*16:data0offset+i*16+4])
	}
}
