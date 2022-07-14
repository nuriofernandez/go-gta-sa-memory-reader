package gtaMemoryAccessor

import (
	. "github.com/0xrawsec/golang-win32/win32"
	"github.com/xXNurioXx/go-gta-sa-memory-reader/memoryReader"
)

func (g *GtaApi) GetPed() DWORD {
	return memoryReader.ReadDWORD(g.PID, 0xB6F5F0)
}
