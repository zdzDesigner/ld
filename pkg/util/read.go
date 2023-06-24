package util

import (
	"bytes"
	"encoding/binary"
)

// 读取内容到结构体
func Read[T any](data []byte) (val T) {

  // 创建可读数据源
	reader := bytes.NewReader(data)
  // 从数据源中以二进制方式读取到T泛型数据中
	if err := binary.Read(reader, binary.LittleEndian, &val); err != nil {
		MustNoErr(err)
	}
	return
}
