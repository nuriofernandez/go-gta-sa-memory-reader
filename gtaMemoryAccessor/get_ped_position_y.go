package gtaMemoryAccessor

import "github.com/xXNurioXx/go-memory-accessor"

func (g *GtaApi) GetPedPositionY() float32 {
	pedPosition := g.GetPedPosition()
	return memoryAccessor.ReadFloat(g.PID, pedPosition+0x34)
}
