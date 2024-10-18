package main

import (
	"fmt"
	"github.com/pbnjay/memory"
)

func getTotalSystemMemory() string {
	TotalSystemMemory := memory.TotalMemory()
	TotalSystemMemoryString := fmt.Sprintf("%d", TotalSystemMemory)
	return TotalSystemMemoryString
}
