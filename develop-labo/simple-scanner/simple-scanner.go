package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

var PortNum int
var MaxNum int = 49151
var Dist string = "ns.ie.u-ryukyu.ac.jp"
var DetectPort string

func PrintPort(ch <-chan string){
		for{
			OpenPort := <-ch
			fmt.Println(OpenPort)
	}
	}

func main() {

	ch := make(chan string)

	go PrintPort(ch)

	PortNum = 1
	for PortNum <= MaxNum {
		//fmt.Println("Scanning port is : ", strconv.Itoa(PortNum))
		_, err := net.DialTimeout("tcp", Dist+":"+strconv.Itoa(PortNum), (150000000 * time.Nanosecond))

		if err != nil {
			//fmt.Println(err)
		} else {
				DetectPort = fmt.Sprintf("%s : port already opend!!!\n",strconv.Itoa(PortNum))
				ch <- DetectPort
			//fmt.Println(reflect.TypeOf(con))
		}
		PortNum = PortNum + 1
	}
}
