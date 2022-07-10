package memoryReader

import "golang.org/x/sys/windows"

func FindProcessByName(processName string) *windows.ProcessEntry32 {
	h, e := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if e != nil {
		panic(e)
	}

	const processEntrySize = 568
	p := windows.ProcessEntry32{Size: processEntrySize}

	for {
		e := windows.Process32Next(h, &p)
		if e != nil {
			break
		}
		s := windows.UTF16ToString(p.ExeFile[:])

		if s == processName {
			println("Found!")
			return &p
		}
	}

	return nil
}
