package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

func readArguments(args []string) (int64, int64) {
	percentageLoad, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		fmt.Println("Invalid percentage load")
		percentageLoad = 0
	}
	secondsToRun, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		fmt.Println("Invalid seconds to run")
		secondsToRun = 0
	}
	return percentageLoad, secondsToRun
}

func calcuatePercentage(percentageLoad int64) float64 {
	return float64((float64(percentageLoad) / 100.0))
}

func calcuateCPUCount(percentage float64, CPU int) float64 {
	return float64(CPU) * percentage
}

func outputLoadAndCount(cpuCount, percentage float64, percentageLoad int64) {
	fmt.Println(fmt.Sprintf("%v * ( %v / %v ) = %v", runtime.NumCPU(), percentageLoad, 100, cpuCount))
	fmt.Println(fmt.Sprintf("%v * %v = %v", runtime.NumCPU(), percentage, cpuCount))
}

func troll(cpuCount float64, secondsToRun int64) {
	done := make(chan int)

	for i := 0; i < int(cpuCount); i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}

	time.Sleep(time.Second * time.Duration(secondsToRun))
	close(done)
}

func main() {
	args := os.Args[1:]
	var percentageLoad, secondsToRun = readArguments(args)
	var percentage = calcuatePercentage(percentageLoad)
	var cpuCount = calcuateCPUCount(percentage, runtime.NumCPU())
	outputLoadAndCount(cpuCount, percentage, percentageLoad)
	troll(cpuCount, secondsToRun)
}
