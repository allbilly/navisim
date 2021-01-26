package rdnaemu

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gitlab.com/akita/mem"
	"gitlab.com/akita/mem/idealmemcontroller"
	"gitlab.com/akita/mem/vm"
	"gitlab.com/akita/navisim/rdnainsts"
	"gitlab.com/akita/util/ca"
)

var _ = Describe("ALU", func() {

	var (
		mockCtrl  *gomock.Controller
		pageTable *MockPageTable

		alu           *ALUImpl
		state         *mockInstState
		storage       *mem.Storage
		addrConverter *idealmemcontroller.InterleavingConverter
		sAccessor     *storageAccessor
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		pageTable = NewMockPageTable(mockCtrl)

		storage = mem.NewStorage(1 * mem.GB)
		addrConverter = &idealmemcontroller.InterleavingConverter{
			InterleavingSize:    1 * mem.GB,
			TotalNumOfElements:  1,
			CurrentElementIndex: 0,
			Offset:              0,
		}
		sAccessor = newStorageAccessor(storage, pageTable, 12, addrConverter)
		alu = NewALU(sAccessor)

		state = new(mockInstState)
		state.scratchpad = make([]byte, 4096)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("should run FLAT_LOAD_USHORT", func() {
		for i := 0; i < 64; i++ {
			pageTable.EXPECT().
				Find(ca.PID(1), uint64(i*4)+4).
				Return(vm.Page{
					PAddr: uint64(0),
				}, true)
		}
		state.inst = rdnainsts.NewInst()
		state.inst.FormatType = rdnainsts.FLAT
		state.inst.Opcode = 10
		state.inst.Offset = rdnainsts.NewIntOperand(10, 4)

		layout := state.Scratchpad().AsFlat()
		for i := 0; i < 64; i++ {
			layout.ADDR[i] = uint64(i * 4)
			storage.Write(uint64(i*4)+4, rdnainsts.Uint32ToBytes(uint32(i)))
		}
		layout.EXEC = 0xffffffffffffffff

		alu.Run(state)

		for i := 0; i < 64; i++ {
			Expect(layout.DST[i*4]).To(Equal(uint32(i)))
			//Expect(layout.DST[i*4+2+4]).To(Equal(uint32(0)))
			//Expect(layout.DST[i*4+3+4]).To(Equal(uint32(0)))
		}
	})

})
