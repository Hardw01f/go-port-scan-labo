package main

import (
	"fmt"
	"net"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var PortNum int
var MaxNum int = 49151
var Dist string = "133.13.50.10"
var DetectPort string

func DebugScan(PortNum int) {
	_, err := net.DialTimeout("tcp", Dist+":"+strconv.Itoa(PortNum), (3000000000 * time.Nanosecond))

	if err != nil {
		fmt.Printf("%d port is closed\n", PortNum)
	} else {
		fmt.Printf("%d  Port Opened\n", PortNum)
	}
}

func Scan(PortNum int, wg *sync.WaitGroup) {
	_, err := net.DialTimeout("tcp", Dist+":"+strconv.Itoa(PortNum), (4000000000 * time.Nanosecond))

	if err != nil {

	} else {
		fmt.Printf("!!!  %d  Port Opened !!!\n", PortNum)
	}
	wg.Done()
}

func main() {

	wg := new(sync.WaitGroup)

	for PortNum = 1; PortNum <= MaxNum; PortNum++ {
		//fmt.Printf("Goroutine : %d\n", runtime.NumGoroutine())
		//fmt.Println(PortNum)
		/*if PortNum == 49151 {
				fmt.Println("last port started")
		}*/

			wg.Add(1)
			go Scan(PortNum, wg)

		fmt.Printf("NumGoroutine : %d\n", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Printf("Use cpu core : %d\n",runtime.NumCPU())
	fmt.Println("finish")
	//DebugScan(5060)
}
