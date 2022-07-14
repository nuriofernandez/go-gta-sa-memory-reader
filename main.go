package main

import (
	"fmt"
	. "github.com/0xrawsec/golang-win32/win32"
	"github.com/xXNurioXx/go-gta-sa-memory-reader/memoryReader"
)

func main() {
	process := memoryReader.FindProcessByName("gta_sa.exe")
	pid := DWORD(process.ProcessID)

	// get ped 0xB6F5F0 *dword
	// get status from ped + 0x46C *uint8
	pedDWORD := memoryReader.ReadDWORD(pid, 0xB6F5F0)
	statusByte := memoryReader.ReadByte(pid, pedDWORD+0x46C)
	money := memoryReader.ReadInt(pid, 0xB7CE50)

	//hexPointer := hex.EncodeToString(baseAddress)
	fmt.Println("Ped status: (0=air/water, 1=in vehicle, 2=entering interior, 3=in floor) ->", statusByte)
	fmt.Println("User money is ->", money)
}
