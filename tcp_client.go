package main

import "net"
import "fmt"
//import "bufio"
//import "os"


// Module is able to receive Hello-message from server when connecting
// Module is able to send message directly to server, and receiving the echo confirming that the
// message has been received by the server


func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}

func main(){
	serverAddr := "129.241.187.23:34933"

	serverConn, err := net.Dial("tcp", serverAddr);
	CheckError(err)

	var listenBuffer[1020]byte
	_, err = serverConn.Read(listenBuffer[0:])

	fmt.Println(string(listenBuffer[0:]))
	
	for {
		_, err = serverConn.Write([]byte("Hello world\x00"))
		CheckError(err)
		_, err = serverConn.Read(listenBuffer[0:])
		CheckError(err)
		fmt.Println(string(listenBuffer[0:]))
	}
}