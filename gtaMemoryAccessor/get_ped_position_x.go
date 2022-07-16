package gtaMemoryAccessor

import "github.com/xXNurioXx/go-memory-accessor"

func (g *GtaApi) GetPedPositionX() float32 {
	pedPosition := g.GetPedPosition()
	return memoryAccessor.ReadFloat(g.PID, pedPosition+0x30)
}
