package gtaMemoryAccessor

import "github.com/xXNurioXx/go-memory-accessor"

func (g *GtaApi) GetPedVehiclePositionY() float32 {
	vehiclePosition := g.GetPedVehiclePosition()
	return memoryAccessor.ReadFloat(g.PID, vehiclePosition+0x34)
}
