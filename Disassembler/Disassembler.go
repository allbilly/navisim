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

	elfFile, err := elf.Open(path)

	if err != nil {
		_ = fmt.Errorf("failed to open file %v", path)
	}
	defer elfFile.Close()

	sec := elfFile.Section(".text")
	data, _ := sec.Data()
	HsaCoData := insts.NewHsaCoFromData(data)

	fmt.Printf("\n%s:\tfile format ELF64-amdgpu\n", filename)
	fmt.Printf("\n\nDisassembly of section .text:\n")

	readform := HsaCoData.InstructionData()
	fmt.Printf("\n\n%v\n\n", readform)

	disasm := insts.NewDisassembler()

	fmt.Printf("\n\n%v\n\n", disasm)

	/*disasm.tryPrintSymbol(*elfFile, sec.Offset, os.Stdout)

	path := os.Args[1]
	elfFile, err := elf.Open(path)
	if err != nil {
		_ = fmt.Errorf("failed to open file %v", path)
	}
	defer elfFile.Close()

	_, filename := filepath.Split(path)

	disasm := insts.NewDisassembler()
	disasm.Disassemble(elfFile, filename, os.Stdout)*/
}
