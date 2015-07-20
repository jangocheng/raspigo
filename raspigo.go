package raspigo

import (
	"strings"
	"strconv"
	"log"
	"os/exec"
)

func byteToFloat(byteIn []byte) float64 {
	stringVal := strings.Replace(string(byteIn), "\n", "", -1)
	temp, err := strconv.ParseFloat(stringVal, 64)
	if err != nil {
		log.Fatal(err)
	}
	return temp
}

func GetCpuTemperature() float64 {
	outByte, err := exec.Command("sh", "-c", "cat /sys/class/thermal/thermal_zone*/temp").Output()
	if err != nil {
		log.Fatal(err)
	}
	return byteToFloat(outByte) / 1000.0
}

func GetCpuVoltage() float64 {
	outByte, err := exec.Command("sh", "-c", "/opt/vc/bin/vcgencmd measure_volts | tr -d 'volt=' | tr -d 'V'").Output()
	if err != nil {
		log.Fatal(err)
	}
	return byteToFloat(outByte)
}

func GetCpuClockSpeed() int {
	outByte, err := exec.Command("sh", "-c", "/opt/vc/bin/vcgencmd measure_clock arm | tr -d 'frequency(45)='").Output()
	if err != nil {
		log.Fatal(err)
	}
	outString := strings.Replace(string(outByte), "\n", "", -1)
	outInt, err := strconv.Atoi(outString)
	return outInt
}