package main

import "fmt"
//import "bytes"
//import "strings"
//import "unicode/utf8"

type somestruct struct {
	data string
	an_int int
	second_int int
}



func main() {

	j := []byte(9)
	fmt.Println(j)
}

/*



func struct_to_bytes(the_struct somestruct) ([]byte){
	length := len(the_struct.data) + len(the_struct.an_int) + len(the_struct.second_int)
	fmt.Println(length)
	result := make([]byte, length)
	fmt.Println(utf8.RuneCountInString(text))
	for i := 0; i < len(the_struct); i++ {
		result[i] = the_struct[i]
	}
	fmt.Println(result)
	return result
}
*/


func string_to_byte_slice(text string) ([]byte){
	result := make([]byte, utf8.RuneCountInString(text), utf8.RuneCountInString(text))
	fmt.Println(utf8.RuneCountInString(text))
	for i := 0; i < len(text); i++ {
		result[i] = text[i]
	}
	fmt.Println(result)
	return result
}