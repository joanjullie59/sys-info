package main

import (
	"fmt"
	"os"
)

func getPIDOfCurrentProcess() string {
	PIDOfCurrentProcess := os.Getpid()
	PIDOfCurrentProcessString := fmt.Sprintf("%d", PIDOfCurrentProcess)
	return PIDOfCurrentProcessString
}
