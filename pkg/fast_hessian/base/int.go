package base

import (
	"encoding/binary"
)

func readInt16(buf []byte) (value int16, length int, err error) {
	
	value = int16(binary.BigEndian.Uint16(buf))
	return value, 2, err
}

func ReadI32(buf []byte) (value int32, length int, err error) {

	value = int32(binary.BigEndian.Uint32(buf))
	return value, 4, err
}

func ReadI64(buf []byte) (value int64, length int, err error) {

	value = int64(binary.BigEndian.Uint64(buf))
	return value, 8, err
}
