package gtaSa

import (
	"fmt"
	. "github.com/0xrawsec/golang-win32/win32"
	"github.com/xXNurioXx/go-gta-sa-memory-reader/gtaMemoryAccessor"
	memoryAccessor "github.com/xXNurioXx/go-memory-accessor"
)

func (gta *GtaSa) Hook() {
	fmt.Println("Trying to find a running GTA:SA instance.")

	process := memoryAccessor.FindProcessByName("gta_sa.exe")

	if process == nil {
		fmt.Println("Unable to find a running GTA:SA instance.")
		return
	}

	pid := DWORD(process.ProcessID)
	gta.GtaApi = gtaMemoryAccessor.GtaApi{PID: pid}

	fmt.Println("Successfully hooked into a running GTA:SA instance.")
}
