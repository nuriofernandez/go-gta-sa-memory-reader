package gtaMemoryAccessor

import (
	"github.com/xXNurioXx/go-gta-sa-memory-reader/memoryReader"
)

func (g *GtaApi) GetPedStatus() byte {
	ped := g.GetPed()
	return memoryReader.ReadByte(g.PID, ped+0x46C)
}
