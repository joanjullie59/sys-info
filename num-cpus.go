package main

import (
	"fmt"
	"runtime"
)

func getCPUCount() string {
	CPUCount := runtime.NumCPU()
	CPUCountString := fmt.Sprintf("%d", CPUCount)
	return CPUCountString
}
