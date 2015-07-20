package main

import (
	"fmt"
	"github.com/xam090/raspigo"
)

func main() {
	cpuTemp := raspigo.GetCpuTemperature()
	fmt.Println("CPU temp:",cpuTemp)
}
