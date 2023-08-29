package fast_hessian

import "fmt"

type DateType byte

const (
	BOOL   = 0
	BYTE   = 1
	I08    = 1
	DOUBLE = 2
	I16    = 3
	I32    = 4
	I64    = 5
)

var typeNames = map[int]string{
	BOOL:   "BOOL",
	BYTE:   "BYTE",
	DOUBLE: "DOUBLE",
	I16:    "I16",
	I32:    "I32",
	I64:    "I64",
}

var typeLengths = map[int]int{
	BOOL:   1,
	BYTE:   1,
	DOUBLE: 8,
	I16:    2,
	I32:    4,
	I64:    8,
}

func (p DateType) String() string {
	if s, ok := typeNames[int(p)]; ok {
		return s
	}
	return "Unknown"
}
func (p DateType) Length() int {
	if s, ok := typeLengths[int(p)]; ok {
		return s
	}
	return 0
}

func (p DateType) CheckBufLength(buf []byte) error {
	dateLength := typeLengths[int(p)]
	if dateLength > len(buf) {
		return fmt.Errorf("[%s] buf length less than %d", typeNames[int(p)], dateLength)
	}
	return nil
}
