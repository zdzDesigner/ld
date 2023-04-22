package linker

import (
	"bytes"
	"ld/pkg/util"
)

func MustELFMagic(contents []byte) {
	if !bytes.HasPrefix(contents, []byte{'\x7f', 'E', 'L', 'F'}) {
		util.FailExit("no ELF file!")
	}

}
