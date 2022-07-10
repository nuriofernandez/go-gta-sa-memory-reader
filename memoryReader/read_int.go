package memoryReader

import (
	. "github.com/0xrawsec/golang-win32/win32"
)

func ReadByte(pid DWORD, pointer DWORD) byte {
	mem := make([]byte, 1)
	read, _ := memoryRead(pid, pointer, mem)
	return read[0]
}
