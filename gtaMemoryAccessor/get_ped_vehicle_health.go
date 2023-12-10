package gtaMemoryAccessor

import (
	"github.com/xXNurioXx/go-memory-accessor"
)

func (g *GtaApi) GetPedVehicleHealth() float32 {
	pedVehicle := g.GetPedVehicle()
	return memoryAccessor.ReadFloat(g.PID, pedVehicle+0x4C0)
}
