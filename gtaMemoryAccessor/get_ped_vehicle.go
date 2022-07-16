package gtaMemoryAccessor

import (
	. "github.com/0xrawsec/golang-win32/win32"
	"github.com/xXNurioXx/go-memory-accessor"
)

func (g *GtaApi) GetPedVehicle() DWORD {
	ped := g.GetPed()
	return memoryAccessor.ReadDWORD(g.PID, ped+0x58C)
}
