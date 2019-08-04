package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

func main() {
	port := 1989
	host := "127.0.0.1"
	conn, err := net.Dial("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal("Error dialing", err.Error())
	}
	// fmt.Fprintf(conn, "hello")
	// fmt.Fprintf(conn, "x")
	buf := make([]byte, 2048)
	for {
		n, _ := conn.Read(buf)
		fmt.Println(n)
		up := make([]byte, 1)
		up[0] = 65
		conn.Write([]byte(up))
		// fmt.Println(buf)
	}
}
