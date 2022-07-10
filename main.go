package main

import (
	"GTASAGO/memoryReader"
	"fmt"
	. "github.com/0xrawsec/golang-win32/win32"
)

func main() {
	process := memoryReader.FindProcessByName("gta_sa.exe")
	pid := DWORD(process.ProcessID)

	// get ped 0xB6F5F0 *dword
	// get status from ped + 0x46C *uint8
	pedDWORD := memoryReader.ReadDWORD(pid, 0xB6F5F0)
	statusByte := memoryReader.ReadByte(pid, pedDWORD+0x46C)

	//hexPointer := hex.EncodeToString(baseAddress)
	fmt.Println("Ped status: (0=air/water, 1=in vehicle, 2=entering interior, 3=in floor) -> ", statusByte)
}
