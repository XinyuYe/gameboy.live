package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

func out() []byte {
	res := "\033[2J\033[H"
	res += "Hello\r\n"
	return []byte(res)
}

func main() {
	port := 5000
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal("Error listening", err.Error())
	}
	conn, err := listener.Accept()
	// Send telnet options
	_, err = conn.Write([]byte{255, 253, 34})
	_, err = conn.Write([]byte{255, 250, 34, 1, 0, 255, 240})
	// NOT ECHO
	_, err = conn.Write([]byte{0xFF, 0xFB, 0x01})
	// buf := make([]byte, 512)
	// n, _ := conn.Read(buf)
	// clean
	data := []byte("\033[2J\033[H")
	_, err = conn.Write(data)
	fmt.Println(data)
	buf := make([]byte, 512)
	n, _ := conn.Read(buf)
	fmt.Println(n)
	fmt.Println(buf)
	n, _ = conn.Read(buf)
	fmt.Println(n)
	fmt.Println(buf)
	// for {
	// 	buf := make([]byte, 512)
	// 	n, _ := conn.Read(buf)
	// fmt.Println(n)
	// fmt.Println(buf)
	// conn.Write([]byte{255, 253, 34})
	// }
}
