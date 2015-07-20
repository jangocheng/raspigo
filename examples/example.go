package main

import (
	"fmt"
	"github.com/xam090/raspigo"
)

func main() {
	cpuTemp := raspigo.GetCpuTemperature()
	fmt.Println("CPU temp:",cpuTemp)
	
	cpuVolt := raspigo.GetCpuVoltage()
	fmt.Println("CPU voltage:",cpuVolt)
	
	cpuClockSpeec := raspigo.GetCpuClockSpeed()
	fmt.Println("CPU clock speed:",cpuClockSpeec)
	
	freeMemory := raspigo.GetFreeMemory()
	fmt.Println("Free memory in %:", freeMemory)
}
