package linker

import "unsafe"

const ELFHeaderSize = int(unsafe.Sizeof(ELFHeader{}))
const SectionHeaderSize = int(unsafe.Sizeof(SectionHeader{}))

type ELFHeader struct {
	Ident     [16]uint8
	Type      uint16
	Machine   uint16
	Version   uint32
	Entry     uint64
	PhOff     uint64
	ShOff     uint64 // section header offset
	Flags     uint32
	EhSize    uint16
	PhEntSize uint16
	PhNum     uint16
	ShEntSize uint16
	ShNum     uint16 // section 的数量
	ShStrndx  uint16
}

type SectionHeader struct {
	Name      uint32
	Type      uint32
	Flags     uint64
	Addr      uint64
	Offset    uint64
	Size      uint64 // 如果 ELFHeader.ShNum 为0, 取当前Size
	Link      uint32
	Info      uint32
	AddrAlign uint64
	EntSize   uint64
}

type ProgramHeader struct {
	Type     uint32
	Flags    uint32
	Offset   uint64
	VAddr    uint64
	PAddr    uint64
	FileSize uint64
	MemSize  uint64
	Align    uint64
}
