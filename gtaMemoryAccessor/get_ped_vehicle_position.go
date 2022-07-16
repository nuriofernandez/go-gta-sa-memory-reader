package gtaMemoryAccessor

import (
	. "github.com/0xrawsec/golang-win32/win32"
	"github.com/xXNurioXx/go-memory-accessor"
)

func (g *GtaApi) GetPedVehiclePosition() DWORD {
	pedVehicle := g.GetPedVehicle()
	return memoryAccessor.ReadDWORD(g.PID, pedVehicle+0x14)
}
