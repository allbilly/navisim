// Package spmm include the benchmark of sparse matrix-matrix multiplication
package spmm

import (
	"fmt"
	"gitlab.com/akita/navisim/benchmarks/matrix/csr"
	"gitlab.com/akita/navisim/driver"
	"log"
	"math/rand"

	// embed hsaco files
	_ "embed"

	"gitlab.com/akita/navisim/kernels"
	"gitlab.com/akita/navisim/rdnainsts"
)

// KernelArgs sets up kernel arguments
type KernelArgs struct {
}

// Benchmark sets up test parameters
type Benchmark struct {
	driver           *driver.Driver
	context          *driver.Context
	gpus             []int
	queues           []*driver.CommandQueue
	useUnifiedMemory bool
	spmmKernel       *rdnainsts.HsaCo

	Dim       int32
	Sparsity  float64
	dAValData driver.GPUPtr
	dBMatData driver.GPUPtr
	dAColData driver.GPUPtr
	dARowData driver.GPUPtr
	dOutData  driver.GPUPtr
	numItems  int32
	matB      []float32
	matOut    []float32
	maxval    float32
	sparseMat csr.Matrix
}

func NewBenchmark(driver *driver.Driver) *Benchmark {
	b := new(Benchmark)
	b.driver = driver
	b.context = driver.Init()
	b.loadProgram()
	b.maxval = 10
	return b
}

// SelectGPU selects GPU
func (b *Benchmark) SelectGPU(gpus []int) {
	b.gpus = gpus
}

// SetUnifiedMemory uses Unified Memory
func (b *Benchmark) SetUnifiedMemory() {
	b.useUnifiedMemory = true
}

//go:embed spmm.hsaco
var hsacoBytes []byte

func (b *Benchmark) loadProgram() {
	b.spmmKernel = kernels.LoadProgramFromMemory(
		hsacoBytes, "spmm_csr_cwm_kernel")
	if b.spmmKernel == nil {
		log.Panic("Failed to load kernel binary")
	}
}

// Run runs the benchmark
func (b *Benchmark) Run() {
	for _, gpu := range b.gpus {
		b.driver.SelectGPU(b.context, gpu)
		b.queues = append(b.queues, b.driver.CreateCommandQueue(b.context))
	}

	b.initMem()
	b.exec()
}

func (b *Benchmark) initMem() {
	b.numItems = int32(float64(b.Dim) * float64(b.Dim) * b.Sparsity)
	fmt.Printf("Number of non-zero elements %d\n", b.numItems)

	b.sparseMat = csr.
		MakeMatrixGenerator(uint32(b.Dim), uint32(b.numItems)).
		GenerateMatrix()
	b.matB = make([]float32, b.Dim*b.Dim)
	b.matOut = make([]float32, b.Dim*b.Dim)

	for i := range b.matB {
		b.matB[i] = rand.Float32() * b.maxval
	}

	var memoryAlloc func(ctx *driver.Context, byteSize uint64) driver.GPUPtr
	if b.useUnifiedMemory {
		memoryAlloc = func(ctx *driver.Context, byteSize uint64) driver.GPUPtr {
			return b.driver.AllocateUnifiedMemory(ctx, byteSize)
		}
	} else {
		memoryAlloc = func(ctx *driver.Context, byteSize uint64) driver.GPUPtr {
			return b.driver.AllocateMemory(ctx, byteSize)
		}
	}
	b.dAValData = memoryAlloc(b.context, uint64(b.numItems*4))
	b.dBMatData = memoryAlloc(b.context, uint64(b.Dim*b.Dim*4))
	b.dAColData = memoryAlloc(b.context, uint64(b.numItems*4))
	b.dARowData = memoryAlloc(b.context, uint64((b.Dim+1)*4))
	b.dOutData = memoryAlloc(b.context, uint64(b.Dim*b.Dim*4))
}

func (b *Benchmark) exec() {
	//TODO: this
}
