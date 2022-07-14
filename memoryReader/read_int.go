package memoryReader

import (
	. "github.com/0xrawsec/golang-win32/win32"
)

func ReadInt(pid DWORD, pointer DWORD) int {
	mem := make([]byte, 4)
	ba, _ := memoryRead(pid, pointer, mem)
	var value int32
	value |= int32(ba[0])
	value |= int32(ba[1]) << 8
	value |= int32(ba[2]) << 16
	value |= int32(ba[3]) << 24
	return int(value)
}
