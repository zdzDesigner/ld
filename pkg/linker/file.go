package linker

import (
	"debug/elf"
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
	// 内容至少要大于头
	if len(f.Contents) < ELFHeaderSize {
		util.FailExit("file too small")
	}

	// 使用泛型把内容读取到结构体ELFHeader中
	elfHdr := util.Read[ELFHeader](f.Contents)
	fmt.Printf("elfHdr:%+v\n", elfHdr)

	// 第一个section header
	firstSectHdr := util.Read[SectionHeader](f.Contents[elfHdr.ShOff:])
	protoSectHdr := protocalSectionHeader(elfHdr, firstSectHdr)

	// 读取section header
	sectHdrs := readSectionHeader(protoSectHdr.ShNum, firstSectHdr, f.Contents[elfHdr.ShOff:])
	util.Assert(len(sectHdrs) == 12, "section header number error")
	fmt.Printf("sectionHeader:%d\n", len(sectHdrs))

	readSectionNames(sectHdrs[protoSectHdr.ShStrndx], f.Contents)

	printSectionHeader(sectHdrs)

}

// ELFHeader 和 SectionHeader 相关协议
func protocalSectionHeader(elfHdr ELFHeader, sectHdr SectionHeader) *ProtocalSectionHeader {
	protoSectHdr := ProtocalSectionHeader{}

	if elfHdr.ShNum == 0 {
		protoSectHdr.ShNum = sectHdr.Size
	} else {
		protoSectHdr.ShNum = uint64(elfHdr.ShNum)
	}
	if int(elfHdr.ShStrndx) == int(elf.SHN_XINDEX) {
		protoSectHdr.ShStrndx = uint64(sectHdr.Link)
	} else {
		protoSectHdr.ShStrndx = uint64(elfHdr.ShStrndx)
	}
	return &protoSectHdr
}

// 读取section header
func readSectionHeader(shnum uint64, firstSectHdr SectionHeader, contents []byte) []SectionHeader {
	sectHdrs := make([]SectionHeader, shnum, shnum)
	sectHdrs[0] = firstSectHdr // 第一个
	i := uint64(1)

	for i < shnum { // 后续section
		contents = contents[SectionHeaderSize:]
		sectHdrs[i] = util.Read[SectionHeader](contents)
		i++
	}
	return sectHdrs
}

func readSectionNames(sectionHeader SectionHeader, contents []byte) []byte {
	return nil

}

func printSectionHeader(sectHdrs []SectionHeader) {
	for _, section := range sectHdrs {
		fmt.Println(section)
	}
}
