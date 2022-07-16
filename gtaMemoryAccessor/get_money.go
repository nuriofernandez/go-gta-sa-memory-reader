package gtaMemoryAccessor

import "github.com/xXNurioXx/go-memory-accessor"

func (g *GtaApi) GetMoney() int {
	return memoryAccessor.ReadInt(g.PID, 0xB7CE50)
}
