package main

import (
	"fmt"
	. "github.com/0xrawsec/golang-win32/win32"
	"github.com/xXNurioXx/go-gta-sa-memory-reader/gtaMemoryAccessor"
	"github.com/xXNurioXx/go-gta-sa-memory-reader/memoryReader"
)

func main() {
	process := memoryReader.FindProcessByName("gta_sa.exe")
	pid := DWORD(process.ProcessID)

	gtaApi := gtaMemoryAccessor.GtaApi{PID: pid}
	statusByte := gtaApi.GetPedStatus()
	money := gtaApi.GetMoney()

	fmt.Println("Ped status: (0=air/water, 1=in vehicle, 2=entering interior, 3=in floor) ->", statusByte)
	fmt.Println("User money is ->", money)
}
