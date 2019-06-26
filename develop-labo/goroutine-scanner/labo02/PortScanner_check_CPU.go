package main

import (
	"fmt"
	"log"
	"net"
	"os"
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

func Scan(PortNum int, semapho chan int, wg *sync.WaitGroup, file *os.File) {
	semapho <- 1

	_, err := net.DialTimeout("tcp", Dist+":"+strconv.Itoa(PortNum), (1500 * time.Millisecond))

	if err != nil {
		/*fmt.Printf(" Checking %d Port\n",PortNum)
		if PortNum%100 ==0{
				fmt.Printf("Goroutine num is       %d\n",runtime.NumGoroutine())
		}*/

	} else {
		fmt.Fprintf(file, "%d Port Opened\n", PortNum)
		fmt.Printf("%d  Port Opened\n", PortNum)
		//fmt.Printf("Goroutine : %d\n", runtime.NumGoroutine())

	}
	<-semapho
	defer wg.Done()
}

func main() {

	wg := new(sync.WaitGroup)
	cpus := runtime.NumCPU()
	fmt.Printf("runtime.CPUS : %d\n", cpus)
	runtime.GOMAXPROCS(cpus)

	var semaphoMultiple int
	if cpus == 2 {
		semaphoMultiple = 150
	} else {
		semaphoMultiple = 300
	}
	fmt.Printf("semaphoMultiple : %d\n", semaphoMultiple)

	semapho := make(chan int, cpus*semaphoMultiple)

	for {
		logfile, err := os.OpenFile("portscan.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()

	jpn, err := time.LoadLocation("Asia/Tokyo")
	nowtime := time.Now().In(jpn)
	fmt.Fprintf(logfile," --- %s --- \n",nowtime)


		for PortNum = 1; PortNum <= MaxNum; PortNum++ {
			//fmt.Printf("Goroutine : %d\n", runtime.NumGoroutine())
			//fmt.Println(PortNum)
			/*if PortNum == Maxnum {
					fmt.Println("last port started")
			}*/

			wg.Add(1)
			go Scan(PortNum, semapho, wg, logfile)
		}
		wg.Wait()
		fmt.Println("finish")
		fmt.Fprintln(logfile,"")
		//DebugScan(5060)
		time.Sleep(180 * time.Second)
	}
}
