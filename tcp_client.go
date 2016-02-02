package main

import "net"
import "fmt"
//import "bufio"
//import "os"

func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}

func main(){
	conn, err := net.Dial("tcp", "129.241.187.23:33546")
	CheckError(err)

	reply := make([]byte, 1024)
	conn.Read(reply)

	fmt.Println("Reply: ", string(reply[0:1024]))
}