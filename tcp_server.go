package main

import "net"
import "fmt"
//import "bytes"
//import "strings"
//import "unicode/utf8"


const (
    CONN_HOST = "129.241.187.161"	// For work station 8
    CONN_PORT = "34933"	// For fixed-length messages
    CONN_TYPE = "tcp"
)

func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}

func main() {

	client_connection := tcp_server_init()
	
	for {
		message :=tcp_server_listen(client_connection)
		fmt.Println(message)
	}
}

func tcp_server_init() (net.Listener){
	serverListener, err := net.Listen("tcp", CONN_HOST+":"+CONN_PORT)
	CheckError(err)
	return serverListener
}


func tcp_server_listen(serverListener net.Listener) (string){

	serverConnection, err := serverListener.Accept();
	CheckError(err)

	welcome_message := []byte("Welcome to this server.")
	_, err = serverConnection.Write(welcome_message)
	CheckError(err)

	message := readMessage(serverConnection)
	serverConnection.Close()

	return message
}



func readMessage(connection net.Conn) (string){	
	incoming_data := make([]byte, 1024)
	
	received_message_length, err := connection.Read(incoming_data)
	CheckError(err)
	
	/*
	if(received_message_length == 0){
		return 0
	}
	*/

	fmt.Println(received_message_length)

	received_message := string(incoming_data[:received_message_length - 1])

	fmt.Println("Incoming message: ", received_message)

	message := []byte("Message received.")

	_, err = connection.Write(message)
	CheckError(err)
	
	return received_message
}