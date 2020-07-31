package main

import (
	"debug/elf"
	"fmt"
	"os"
	"path/filepath"

	"gitlab.com/akita/mgpusim/insts"
)

func main() {
	path := os.Args[1]

	_, filename := filepath.Split(path)

	fmt.Printf("You opened file from %v\n", filename)

	elfFile, err := elf.Open(path)

	if err != nil {
		_ = fmt.Errorf("failed to open file %v", path)
	}
	defer elfFile.Close()

	sec := elfFile.Section(".text")
	data, _ := sec.Data()
	HsaCoData := insts.NewHsaCoFromData(data)

	readform := HsaCoData.InstructionData()
	fmt.Printf("\n\n%v\n\n", readform)
}
