package main

import (
	"fmt"
)

// 1. Available  Memory
// 2. Total System Memory
// 3. number of CPU
// 4. number of running processes
// 5. the PID of the current process
// 6. IPV4 of the computer

func main() {
	fmt.Println("getFreeMemory()", getFreeMemory())
	fmt.Println("getTotalSystemMemory()", getTotalSystemMemory())
	fmt.Println("getCPUCount()", getCPUCount())
	fmt.Println("getNumberOfProcesses()", getNumberOfProcesses())
	fmt.Println("getPIDOfCurrentProcess()", getPIDOfCurrentProcess())
	fmt.Println("getLocalIP()", getLocalIP())
}
