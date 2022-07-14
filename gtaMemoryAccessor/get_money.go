package gtaMemoryAccessor

import (
	"github.com/xXNurioXx/go-gta-sa-memory-reader/memoryReader"
)

func (g *GtaApi) GetMoney() int {
	return memoryReader.ReadInt(g.PID, 0xB7CE50)
}
