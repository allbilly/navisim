package vectoradd

import (
	"testing"

	"gitlab.com/akita/akita"
	"gitlab.com/akita/navisim/driver"
	"gitlab.com/akita/navisim/platform"
)

func TestAddSignedInputs(t *testing.T) {
	platforms := []struct {
		name  string
		build func() (akita.Engine, *driver.Driver)
	}{
		{"emulation", platform.MakeEmuBuilder().WithNumGPU(1).Build},
		{"timing", platform.MakeR9NanoBuilder().WithNumGPU(1).WithoutProgressBar().Build},
	}

	for _, p := range platforms {
		t.Run(p.name, func(t *testing.T) {
			_, gpuDriver := p.build()
			gpuDriver.Run()
			defer gpuDriver.Terminate()

			input := []float32{-2.5, -1.5, -1, 0, 1, 2.5}
			b := NewBenchmark(gpuDriver)
			b.SelectGPU([]int{1})
			b.Length = len(input)
			b.initMem()
			copy(b.inputData, input)
			gpuDriver.MemCopyH2D(b.context, b.gInputData, b.inputData)
			b.exec()

			for i, value := range input {
				if got, want := b.outputData[i], value+1; got != want {
					t.Errorf("output[%d] = %v, want %v", i, got, want)
				}
			}
		})
	}
}
