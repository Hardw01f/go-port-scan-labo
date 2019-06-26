package main

import (
	"fmt"
	"runtime"
)

func main() {
		cpus := runtime.NumCPU()
		fmt.Printf("runtime.NumCPU vuleu : %d\n",cpus)

		fmt.Printf("runtime.NumCPU direct : %d\n",runtime.NumCPU)
}
