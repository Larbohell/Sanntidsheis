// Go 1.2
// go run helloworld_go.go

package main

import (
    . "fmt"
    "runtime"
    "time"
)

var i int = 0

func thread_1() {
	for j:=0; j<1000000; j++{
		i++
	}
}

func thread_2(){

	for j:=0; j<1000000; j++{
		i--
	}
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
                                            
	go thread_1()
	go thread_2()

    time.Sleep(100*time.Millisecond)
    Println(i)
}
