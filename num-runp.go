package main

import "golang.org/x/sys/windows"

const processEntrySize = 568

func getNumberOfProcesses() int {
	h, e := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if e != nil {
		return 0
	}

	p := windows.ProcessEntry32{Size: processEntrySize}
	processCount := 0

	for {
		e := windows.Process32Next(h, &p)
		if e != nil {
			break
		}
		processCount++
	}

	return processCount
}
