package main

import "net"
import "fmt"
import "time"

func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}


func main(){
	//ListenerAddr, err := net.ResolveUDPAddr("udp4", ":30000");
	//CheckError(err)


	// Acquiring server address
	/*
	serverConn, err := net.ListenUDP("udp4", ListenerAddr);
	CheckError(err)
	buffer := make([]byte,1024)
	size, _, err := serverConn.ReadFromUDP(buffer);
	CheckError(err)
	fmt.Println("Received ", string(buffer[0:size]))
	//Server address acquired
	*/

	// Establish connection for sending
	//ServerAddr, err := net.ResolveUDPAddr("udp", "129.241.287.23:52038");
	//CheckError(err)
	

	broadcastAddr, err := net.ResolveUDPAddr("udp4", "129.241.187.23:20011")
	CheckError(err)
	//conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)

	//LocalAddr, err := net.ResolveUDPAddr("udp4", "129.241.187.159:20011")
	//CheckError(err)


	broadcastConn, err := net.DialUDP("udp4", nil, broadcastAddr)
	CheckError(err)

	
	//listenConn, err := net.ListenUDP("udp4", ListenerAddr)
	//buf := make([]byte, 1024)
	//size, _, err = connListen.ReadFromUDP(buf);

	//CheckError(err)
	//fmt.Println("Received ", string(buffer[0:size]))

	for {

		_ , err := broadcastConn.Write([]byte("hello world\x00"))
		CheckError(err)

		var readBuffer [1024]byte
		_, _, error := broadcastConn.ReadFromUDP(readBuffer[0:])
		CheckError(error)


		//size, _, error := listenConn.ReadFromUDP(buf);
		//CheckError(error)
		fmt.Println("Received ", string(readBuffer[0:1024]))
		
		time.Sleep(time.Second * 1)

		
	}
}

//UDP server at 129.241.287.23:52038
