package memoryReader

import (
	. "github.com/0xrawsec/golang-win32/win32"
	"github.com/0xrawsec/golang-win32/win32/kernel32"
	"golang.org/x/sys/windows"
)

func memoryRead(pid DWORD, pointer DWORD, buff []byte) ([]byte, bool) {
	win32handle, _ := kernel32.OpenProcess(0x0010|windows.PROCESS_VM_READ|windows.PROCESS_QUERY_INFORMATION, BOOL(0), pid)
	_, err := kernel32.ReadProcessMemory(win32handle, LPCVOID(pointer), buff)
	if err != nil {
		return buff, false
	}
	return buff, true
}
