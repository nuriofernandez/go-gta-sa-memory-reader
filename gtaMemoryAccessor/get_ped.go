package gtaMemoryAccessor

import (
	. "github.com/0xrawsec/golang-win32/win32"
	"github.com/xXNurioXx/go-memory-accessor"
)

func (g *GtaApi) GetPed() DWORD {
	return memoryAccessor.ReadDWORD(g.PID, 0xB6F5F0)
}
