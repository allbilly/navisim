package main

import (
	"flag"

	"gitlab.com/akita/navisim/benchmarks/dnn/vectoradd"
	"gitlab.com/akita/navisim/samples/runner"
)

var length = flag.Int("length", 4096, "Number of elements to add.")

func main() {
	flag.Parse()

	runner := new(runner.Runner).ParseFlag().Init()

	benchmark := vectoradd.NewBenchmark(runner.GPUDriver)
	benchmark.Length = *length

	runner.AddBenchmark(benchmark)

	runner.Run()
}
