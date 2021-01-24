package rdnaemu

import (
	"log"

	"gitlab.com/akita/navisim/insts"
	"gitlab.com/akita/navisim/rdnainsts"
)

//nolint:gocyclo
//nolint:funlen
func (u *ALUImpl) runFlat(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {
	case 12:
		u.runFlatLoadDWord(state)
	case 14:
		u.runFlatLoadDwordx4(state)
	case 28:
		u.runFlatStoreDWord(state)
	default:
		log.Panicf("Opcode %d for FLAT format is not implemented", inst.Opcode)
	}
}
func (u *ALUImpl) runFlatLoadDWord(state InstEmuState) {
	sp := state.Scratchpad().AsFlat()
	pid := state.PID()
	for i := uint(0); i < 64; i++ {
		if !laneMasked(sp.EXEC, i) {
			continue
		}

		buf := u.storageAccessor.Read(pid, sp.ADDR[i], uint64(4))
		sp.DST[i*4] = rdnainsts.BytesToUint32(buf)
	}
}

func (u *ALUImpl) runFlatLoadDwordx4(state InstEmuState) {
	sp := state.Scratchpad().AsFlat()
	pid := state.PID()
	for i := uint(0); i < 64; i++ {
		if !laneMasked(sp.EXEC, i) {
			continue
		}

		buf := u.storageAccessor.Read(pid, sp.ADDR[i], uint64(16))

		sp.DST[i*4] = insts.BytesToUint32(buf[0:4])
		sp.DST[i*4+1] = insts.BytesToUint32(buf[4:8])
		sp.DST[i*4+2] = insts.BytesToUint32(buf[8:12])
		sp.DST[i*4+3] = insts.BytesToUint32(buf[12:16])
	}
}

func (u *ALUImpl) runFlatStoreDWord(state InstEmuState) {
	sp := state.Scratchpad().AsFlat()
	pid := state.PID()

	for i := uint(0); i < 64; i++ {
		if !laneMasked(sp.EXEC, i) {
			continue
		}

		u.storageAccessor.Write(
			pid, sp.ADDR[i], rdnainsts.Uint32ToBytes(sp.DATA[i*4]))
	}
}
