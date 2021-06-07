// Package spmm include the benchmark of sparse matrix-matrix multiplication
package spmm

import (
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
	spmmKernel *rdnainsts.HsaCo
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
