package memoryReader

import (
	"encoding/binary"
	. "github.com/0xrawsec/golang-win32/win32"
	"math"
)

func ReadFloat(pid DWORD, pointer DWORD) float32 {
	mem := make([]byte, 4)
	read, _ := memoryRead(pid, pointer, mem)
	bits := binary.LittleEndian.Uint32(read)
	return math.Float32frombits(bits)
}
