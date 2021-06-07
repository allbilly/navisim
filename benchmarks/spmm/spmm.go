// Package spmm include the benchmark of sparse matrix-matrix multiplication
package spmm

import (
	"gitlab.com/akita/navisim/driver"
	"log"

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
	spmmKernel *rdnainsts.HsaCo

	maxval    float32
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
	//TODO: this
}

func (b *Benchmark) exec() {
	//TODO: this
}