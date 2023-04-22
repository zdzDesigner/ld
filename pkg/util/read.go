package util

import (
	"bytes"
	"encoding/binary"
)

// 读取内容到结构体
func Read[T any](data []byte) (val T) {

	reader := bytes.NewReader(data)
	if err := binary.Read(reader, binary.LittleEndian, &val); err != nil {
		MustNoErr(err)
	}
	return
}
