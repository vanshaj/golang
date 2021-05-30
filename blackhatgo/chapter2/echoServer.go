package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:20081")
	if err != nil {
		fmt.Println("Not listening ", err)
	}
	for {
		fmt.Println("Listening on port 20081")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Not listening ", err)
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()
	b := make([]byte, 1024)
	for {
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			fmt.Println("Connection disconnectd from client")
			break
		}
		if err != nil {
			fmt.Println("unexpected ", err)
		}
		_, err = conn.Write(b[0:size])
		if err != nil {
			fmt.Println("Unable to write", err)
		}
	}
}
