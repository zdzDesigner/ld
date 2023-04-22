package linker

import (
	"fmt"
	"ld/pkg/util"
	"os"
)

type File struct {
	Name     string
	Contents []byte
}

func MustNewFile(filename string) *File {
	contents, err := os.ReadFile(filename)
	util.MustNoErr(err)
	return &File{
		Name:     filename,
		Contents: contents,
	}
}

func ReadFile(filename string) {

	f := MustNewFile(filename)
	MustELFMagic(f.Contents)
	if len(f.Contents) < ELFHeaderSize {
		util.FailExit("file too small")
	}

	elfHeader := util.Read[ELFHeader](f.Contents)

	sectionHeaders := readSectionHeader(uint64(elfHeader.ShNum), f.Contents[elfHeader.ShOff:])

	fmt.Printf("elfHeader:%+v\n", elfHeader)
	fmt.Printf("sectionHeader:%d\n", len(sectionHeaders))

}

func readSectionHeader(shnum uint64, contents []byte) []SectionHeader {
	firstSectionHeader := util.Read[SectionHeader](contents)
	// elf shnum 协议
	if shnum == 0 {
		shnum = firstSectionHeader.Size
	}
	sectionHeaders := make([]SectionHeader, shnum, shnum)
	sectionHeaders[0] = firstSectionHeader // 第一个
	i := uint64(1)

	for i < shnum { // 后续section
		contents = contents[SectionHeaderSize:]
		sectionHeaders[i] = util.Read[SectionHeader](contents)
		i++
	}
	return sectionHeaders
}
