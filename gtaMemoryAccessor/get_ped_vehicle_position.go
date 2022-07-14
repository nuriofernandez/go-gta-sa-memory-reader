package gtaMemoryAccessor

import (
	. "github.com/0xrawsec/golang-win32/win32"
	"github.com/xXNurioXx/go-gta-sa-memory-reader/memoryReader"
)

func (g *GtaApi) GetPedVehiclePosition() DWORD {
	pedVehicle := g.GetPedVehicle()
	return memoryReader.ReadDWORD(g.PID, pedVehicle+0x14)
}
