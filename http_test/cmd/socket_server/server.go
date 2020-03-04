package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func tcp() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go NewConn(conn)
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
func NewConn(c net.Conn) {
	buff := make([]byte, 128)
	n, err := c.Read(buff)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("read:", n)
	c.Write([]byte(strings.Repeat(string(buff), 100)))
	c.Close()
}

func udp() {
	add, err := net.ResolveUDPAddr("tcp", "localhost:3000")
	checkError(err)
	conn, err := net.ListenUDP("udp", add)
	checkError(err)
	for {
		fmt.Println("remote:", conn.RemoteAddr().String())
		conn.ReadMsgUDP("")
	}
}
