package memoryReader

import (
	"encoding/binary"
	. "github.com/0xrawsec/golang-win32/win32"
)

func ReadDWORD(pid DWORD, pointer DWORD) DWORD {
	var mem = make([]byte, 4)
	byteArray, _ := memoryRead(pid, pointer, mem)
	return DWORD(binary.LittleEndian.Uint32(byteArray))
}
