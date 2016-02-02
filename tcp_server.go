package main

import "net"
import "fmt"
import "bufio"
//import "strings"



func main() {
	fmt.Print("1\n")

	listener, _:= net.Listen("tcp", "127.0.0.1:33546")
	//if (error){}
	fmt.Print("2\n")

	connection, _:= listener.Accept()
	//if (error){}
	fmt.Print("3\n")

	for {
		fmt.Print("new for\n")
		message, _ := bufio.NewReader(connection).ReadString('\n')

		fmt.Print("Message Received:", string(message))

	}
}

//uint8 [1024] buffer
//uint32 fromWho
