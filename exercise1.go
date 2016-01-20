// Go 1.2
// go run helloworld_go.go

package main

import (
    . "fmt"
    "runtime"
    "time"
)

var i int = 0

func thread_1(channel chan int, chfin chan bool) {
	for j:=0; j<1000000; j++{
		x := <-channel				
		x++
		channel <- x
		//i = <-channel
		//channel <- i
	}
	chfin <- true
	
}

func thread_2(channel chan int, chfin chan bool){
	
	for j:=0; j<1000000; j++{
		x := <-channel				
		x--
		channel <- x
		//i = <-channel
		//channel <- i
	}
	chfin <- true
}

func main() {
	channel := make(chan int, 1)
	channel <- i

	chfin1 := make(chan bool)
	chfin2 := make(chan bool)
	

    	runtime.GOMAXPROCS(runtime.NumCPU())
                                            
	go thread_1(channel, chfin1)
	go thread_2(channel, chfin2)

    	//time.Sleep(100*time.Millisecond)
	//if chfin
	i = <-channel    	
	Println(i)

}
