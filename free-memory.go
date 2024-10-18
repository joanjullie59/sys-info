package main

import (
	"fmt"
	"github.com/pbnjay/memory"
)

func getFreeMemory() string {
	availableMemory := memory.FreeMemory()
	availableMemoryString := fmt.Sprintf("%d", availableMemory)
	return availableMemoryString
}
