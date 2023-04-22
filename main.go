package main

import (
	"fmt"
	"ld/pkg/linker"
	"ld/pkg/util"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		util.FailExit("args invalid!")
	}

	linker.ReadFile(os.Args[1])

	fmt.Println("init", os.Args, len(os.Args), linker.ELFHeaderSize, linker.SectionHeaderSize)
}
