package main

import "net"
import "fmt"

// Module is able to receive Hello-message from server when connecting
// Module is able to send message directly to server, and receiving the echo confirming that the
// message has been received by the server


func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}

type myStruct struct {
	FLOOR int
	elevator_ID int
}

func main(){
	serverConn := tcp_init("129.241.187.159", "34933")
	tcp_send("hei", serverConn)
	
}

func tcp_client_init(serverIP string, port string) (net.Conn) {
	serverConn, err := net.Dial("tcp", serverIP+":"+port);
	CheckError(err)

	var listenBuffer[1020]byte
	_, err = serverConn.Read(listenBuffer[0:])

	fmt.Println(string(listenBuffer[0:]))
	return serverConn
}

func tcp_send(message string, serverConn net.Conn){
	_, err := serverConn.Write([]byte(message+"\x00"))
	CheckError(err)

	var response[1020]byte
	_, err = serverConn.Read(response[0:])
	CheckError(err)
}