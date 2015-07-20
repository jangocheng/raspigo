package raspigo

import (
	"strings"
	"strconv"
	"log"
	"os/exec"
	"net/http"
	"encoding/json"
	"fmt"
)

type RaspiStat struct {
	CpuTemp float64
	CpuVolt float64
	CpuClock int
	FreeMemP int
}

func byteToFloat(byteIn []byte) float64 {
	stringVal := strings.Replace(string(byteIn), "\n", "", -1)
	floatVal, err := strconv.ParseFloat(stringVal, 64)
	if err != nil {
		log.Fatal(err)
	}
	return floatVal
}
func byteToInt(byteIn []byte) int {
	stringVal := strings.Replace(string(byteIn), "\n", "", -1)
	stringVal = strings.Replace(stringVal, " ", "", -1)
	intVal, err := strconv.Atoi(stringVal)
	if err != nil {
		log.Fatal(err)
	}
	return intVal
}

func GetCpuTemperature() float64 {
	outByte, err := exec.Command("sh", "-c", "vcgencmd measure_temp |egrep -o '[1-9]{1,3}\\.[1-9]'").Output()
	if err != nil {
		log.Fatal(err)
	}
	return byteToFloat(outByte)
}

func GetCpuVoltage() float64 {
	outByte, err := exec.Command("sh", "-c", "vcgencmd measure_volts | egrep -o '[0-9]\\.[0-9]{3}'").Output()
	if err != nil {
		log.Fatal(err)
	}
	return byteToFloat(outByte)
}

func GetCpuClockSpeed() int {
	outByte, err := exec.Command("sh", "-c", "vcgencmd measure_clock arm | tr -d 'frequency(45)='").Output()
	if err != nil {
		log.Fatal(err)
	}
	return byteToInt(outByte)
}

func GetFreeMemory() int {
	outByte, err := exec.Command("sh", "-c", "df | awk '{i=0;if($1 == \"rootfs\"){i=1;}if(i==1){print substr($5, 1, 2);}}'").Output()
	if err != nil {
		log.Fatal(err)
	}
	return byteToInt(outByte)
}
func raspiStats(w http.ResponseWriter, r *http.Request) {
	stat := RaspiStat{CpuTemp: GetCpuTemperature(), CpuVolt: GetCpuVoltage(), 
		CpuClock: GetCpuClockSpeed(), FreeMemP: GetFreeMemory()}
	jsonData, err := json.Marshal(stat)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(jsonData))
}
func GetRaspiStatHandler() http.Handler {
	return http.HandlerFunc(raspiStats)
}
