package raspigo

import (
	"strings"
	"strconv"
	"io/ioutil"
	"log"
)

func GetCpuTemperature() float64 {
	outByte, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		log.Fatal(err)
	}
	outString := strings.Replace(string(outByte), "\n", "", -1)
	temp, err := strconv.ParseFloat(outString, 64)
	if err != nil {
		log.Fatal(err)
	}	
	return temp / 1000.0
}
