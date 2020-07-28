package main

import (
	"debug/elf"
	"fmt"
	"os"

	"gitlab.com/akita/mgpusim/insts"
)

func main() {

	//Fetches the file to be disassembled
	path := os.Args[1]

	fmt.Printf("You opened file from %v\n", path)

	elfFile, err := elf.Open(path)

	//If file is not able to be opened throw an error message
	if err != nil {
		_ = fmt.Errorf("failed to open file %v", path)
	}
	defer elfFile.Close()

	//Gets the data
	sec := elfFile.Section(".text")
	data, _ := sec.Data()
	//Creates a new HsaCo object that is empty Stores data from
	//file into HsaCoData
	HsaCoData := insts.NewHsaCoFromData(data)

	//Prints the HsaCoHeader in a user understandable way.
	readform := HsaCoData.Info()
	fmt.Printf(readform)

}
