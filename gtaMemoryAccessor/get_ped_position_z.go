package gtaMemoryAccessor

import (
	"github.com/xXNurioXx/go-gta-sa-memory-reader/memoryReader"
)

func (g *GtaApi) GetPedPositionZ() float32 {
	pedPosition := g.GetPedPosition()
	return memoryReader.ReadFloat(g.PID, pedPosition+0x38)
}
