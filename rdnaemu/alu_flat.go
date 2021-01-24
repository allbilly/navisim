package rdnaemu

import (
	"log"

	"gitlab.com/akita/navisim/rdnainsts"
)

//nolint:gocyclo
//nolint:funlen
func (u *ALUImpl) runFlat(state InstEmuState) {
	inst := state.Inst()
	switch inst.Opcode {
	case 12:
		u.runFlatLoadDWord(state)
	case 28:
		u.runFlatStoreDWord(state)
	case 13:
		u.runFlatLoadDWordX2(state)
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
func (u *ALUImpl) runFlatLoadDWordX2(state InstEmuState) {
	sp := state.Scratchpad().AsFlat()
	pid := state.PID()
	for i := uint(0); i < 64; i++ {
		if !laneMasked(sp.EXEC, i) {
			continue
		}

		buf := u.storageAccessor.Read(pid, sp.ADDR[i], uint64(8))
		sp.DST[i*4] = rdnainsts.BytesToUint32(buf[0:4])
		sp.DST[i*4+1] = rdnainsts.BytesToUint32(buf[4:8])
	}
}
