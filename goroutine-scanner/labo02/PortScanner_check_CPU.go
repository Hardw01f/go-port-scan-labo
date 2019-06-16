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
var MaxNum int = 65536
var Dist string = "localhost"
var DetectPort string

func DebugScan(PortNum int) {
	_, err := net.DialTimeout("tcp", Dist+":"+strconv.Itoa(PortNum), (500 * time.Millisecond))

	if err != nil {
		fmt.Printf("%d port is closed\n", PortNum)
	} else {
		fmt.Printf("%d  Port Opened\n", PortNum)
	}
}

func Scan(PortNum int, semapho chan int, wg *sync.WaitGroup) {
	semapho <- 1
	_, err := net.DialTimeout("tcp", Dist+":"+strconv.Itoa(PortNum), (500 * time.Millisecond))

	if err != nil {
		/*fmt.Printf(" Checking %d Port\n",PortNum)
		if PortNum%100 ==0{
				fmt.Printf("Goroutine num is       %d\n",runtime.NumGoroutine())
		}*/

	} else {
		fmt.Printf("%d  Port Opened\n", PortNum)
		//fmt.Printf("Goroutine : %d\n", runtime.NumGoroutine())

	}
	<-semapho
	defer wg.Done()
}

func main() {

	for{
	wg := new(sync.WaitGroup)
	cpus := runtime.NumCPU()
	fmt.Printf("runtime.CPUS : %d\n", cpus)
	runtime.GOMAXPROCS(cpus)

	semapho := make(chan int, cpus*300)

	for PortNum = 1; PortNum <= MaxNum; PortNum++ {
		//fmt.Printf("Goroutine : %d\n", runtime.NumGoroutine())
		//fmt.Println(PortNum)
		/*if PortNum == Maxnum {
				fmt.Println("last port started")
		}*/

		wg.Add(1)
		go Scan(PortNum, semapho, wg)
	}
	wg.Wait()
	fmt.Println("finish")
	//DebugScan(5060)
	time.Sleep(600 * time.Second)
	}
}
