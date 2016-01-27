// Go 1.2
// go run helloworld_go.go

package main

import (
	. "fmt"
	//"runtime"
)

func thread_1(channel chan int, workerDoneCh chan int) {
	for j := 0; j < 1000000; j++ {
		channel <- 1
	}
	workerDoneCh <- 1

}

func thread_2(channel chan int, workerDoneCh chan int) {
	for j := 0; j < 1000000; j++ {
		channel <- 1
	}
	workerDoneCh <- 1
}

func main() {
	var i int = 0
	channel1 := make(chan int)
	channel2 := make(chan int)
	workerDoneCh := make(chan int)

	//runtime.GOMAXPROCS(runtime.NumCPU())

	go thread_1(channel1, workerDoneCh)
	go thread_2(channel2, workerDoneCh)


	const numWorkers = 2
	numWorkersDone := 0
	done := false
	for !done {
		select {
		case <-channel1:
			i++
		case <-channel2:
			i--
		case <-workerDoneCh:
			Println("worker done")
			numWorkersDone++
			if numWorkersDone == numWorkers {
				done = true
			}
		}
	}
	Println(i)

}
