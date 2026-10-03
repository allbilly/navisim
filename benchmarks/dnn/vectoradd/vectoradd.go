// Package vectoradd implements element-wise vector addition as a benchmark.
package vectoradd

import (
	"log"

	"gitlab.com/akita/navisim/driver"
	"gitlab.com/akita/navisim/kernels"
	"gitlab.com/akita/navisim/rdnainsts"
)

// KernelArgs defines kernel arguments.
type KernelArgs struct {
	Count               uint32
	Padding             uint32
	Input               driver.GPUPtr
	Output              driver.GPUPtr
	HiddenGlobalOffsetX int64
	HiddenGlobalOffsetY int64
	HiddenGlobalOffsetZ int64
}

// Benchmark defines a vector-add benchmark.
type Benchmark struct {
	driver  *driver.Driver
	context *driver.Context
	gpus    []int
	hsaco   *rdnainsts.HsaCo

	Length      int
	inputData   []float32
	outputData  []float32
	gInputData  driver.GPUPtr
	gOutputData driver.GPUPtr

	useUnifiedMemory bool
}

// NewBenchmark returns a benchmark.
func NewBenchmark(driver *driver.Driver) *Benchmark {
	b := new(Benchmark)

	b.driver = driver
	b.context = driver.Init()

	hsacoBytes := _escFSMustByte(false, "/kernels.hsaco")

	// Patched legacy ReLU hsaco (global_load + v_add_f32 + global_store).
	b.hsaco = kernels.LoadProgramFromMemory(hsacoBytes, "ReLUForward")

	return b
}

// SelectGPU selects GPU.
func (b *Benchmark) SelectGPU(gpus []int) {
	b.gpus = gpus
}

// SetUnifiedMemory uses Unified Memory.
func (b *Benchmark) SetUnifiedMemory() {
	b.useUnifiedMemory = true
}

// Run runs the benchmark.
func (b *Benchmark) Run() {
	// The legacy kernel compares its count argument as a signed 32-bit value.
	if b.Length <= 0 || uint64(b.Length) > uint64(1<<31-1) {
		panic("vectoradd length must be between 1 and 2147483647")
	}
	if len(b.gpus) == 0 {
		panic("vectoradd requires at least one GPU")
	}
	b.driver.SelectGPU(b.context, b.gpus[0])
	b.initMem()
	b.exec()
}

func (b *Benchmark) initMem() {
	byteSize := uint64(b.Length) * 4

	if b.useUnifiedMemory {
		b.gInputData = b.driver.AllocateUnifiedMemory(b.context, byteSize)
		b.gOutputData = b.driver.AllocateUnifiedMemory(b.context, byteSize)
	} else {
		b.gInputData = b.driver.AllocateMemory(b.context, byteSize)
		b.driver.Distribute(b.context, b.gInputData, byteSize, b.gpus)

		b.gOutputData = b.driver.AllocateMemory(b.context, byteSize)
		b.driver.Distribute(b.context, b.gOutputData, byteSize, b.gpus)
	}

	b.inputData = make([]float32, b.Length)
	b.outputData = make([]float32, b.Length)
	for i := 0; i < b.Length; i++ {
		b.inputData[i] = float32(i)
	}

	b.driver.MemCopyH2D(b.context, b.gInputData, b.inputData)
}

func (b *Benchmark) exec() {
	queues := make([]*driver.CommandQueue, len(b.gpus))

	for i, gpu := range b.gpus {
		start := int(uint64(b.Length) * uint64(i) / uint64(len(b.gpus)))
		end := int(uint64(b.Length) * uint64(i+1) / uint64(len(b.gpus)))
		if start == end {
			continue
		}

		b.driver.SelectGPU(b.context, gpu)
		q := b.driver.CreateCommandQueue(b.context)
		queues[i] = q

		numWI := end - start

		kernArg := KernelArgs{
			// Bound each GPU's final wavefront to its assigned global indices.
			uint32(end), 0,
			b.gInputData, b.gOutputData,
			int64(start), 0, 0,
		}

		b.driver.EnqueueLaunchKernel(
			q,
			b.hsaco,
			[3]uint32{uint32(numWI), 1, 1},
			[3]uint16{64, 1, 1},
			&kernArg,
		)
	}

	for _, q := range queues {
		if q != nil {
			b.driver.DrainCommandQueue(q)
		}
	}

	b.driver.MemCopyD2H(b.context, b.outputData, b.gOutputData)
}

// Verify verifies the result.
func (b *Benchmark) Verify() {
	for i := 0; i < b.Length; i++ {
		expected := b.inputData[i] + 1
		if b.outputData[i] != expected {
			log.Panicf("mismatch at %d, expected %f, got %f",
				i, expected, b.outputData[i])
		}
	}

	log.Printf("Passed!\n")
}
