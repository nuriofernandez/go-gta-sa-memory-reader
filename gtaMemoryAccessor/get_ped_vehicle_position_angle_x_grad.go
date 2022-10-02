package gtaMemoryAccessor

import (
	"github.com/xXNurioXx/go-memory-accessor"
)

func (g *GtaApi) GetPedVehiclePositionAngleXGrad() float32 {
	pedVehicle := g.GetPedVehiclePosition()
	return memoryAccessor.ReadFloat(g.PID, pedVehicle+0)
}
