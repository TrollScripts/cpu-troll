package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {

	args := os.Args[1:]

	percentageLoad, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		fmt.Println("Invalid percentage load")
	}

	secondsToRun, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		fmt.Println("Invalid seconds to run")
	}

	var percentage float64 = float64((float64(percentageLoad) / 100.0))
	var cpuCount float64 = float64(runtime.NumCPU()) * percentage
	fmt.Println(fmt.Sprintf("%v * ( %v / %v ) = %v", runtime.NumCPU(), percentageLoad, 100, cpuCount))
	fmt.Println(fmt.Sprintf("%v * %v = %v", runtime.NumCPU(), percentage, cpuCount))

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
