package main

import "net"
import "fmt"
//import "bytes"
//import "strings"
import "unicode/utf8"


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

	serverListener, err := net.Listen("tcp", CONN_HOST+":"+CONN_PORT)
	CheckError(err)

	serverConnection, err := serverListener.Accept();
	CheckError(err)

	//welcome_message := []byte("Welcome to this server.")
	welcome_message := []byte("Welcome to this server.")
	_, err = serverConnection.Write(welcome_message)
	CheckError(err)

	for {
		_ = handleRequest(serverConnection)
	}

}

func handleRequest(connection net.Conn) (int){	
	incoming_data := make([]byte, 1024)
	
	received_message_length, err := connection.Read(incoming_data)
	CheckError(err)

	received_message := string(incoming_data[:received_message_length - 1])

	fmt.Println("Incoming message: ", received_message)

	message := []byte("Message received.")

	_, err = connection.Write(message)
	CheckError(err)
	
	return 1
}