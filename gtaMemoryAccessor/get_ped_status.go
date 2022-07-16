package gtaMemoryAccessor

import "github.com/xXNurioXx/go-memory-accessor"

func (g *GtaApi) GetPedStatus() byte {
	ped := g.GetPed()
	return memoryAccessor.ReadByte(g.PID, ped+0x46C)
}
