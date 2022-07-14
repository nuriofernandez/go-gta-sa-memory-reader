package gtaMemoryAccessor

import (
	. "github.com/0xrawsec/golang-win32/win32"
	"github.com/xXNurioXx/go-gta-sa-memory-reader/memoryReader"
)

func (g *GtaApi) GetPedPosition() DWORD {
	ped := g.GetPed()
	return memoryReader.ReadDWORD(g.PID, ped+0x14)
}
