package main

import (
	"io"
	"log"
	"net"
)

func handle(conn net.Conn) {
	destConn, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err != nil {
		log.Fatalln("Unable to connect to scan me ", err)
	}
	go func() {
		_, err = io.Copy(destConn, conn)
		if err != nil {
			log.Fatalln("Unable to copy to dest writer ", err)
		}
	}()

	_, err = io.Copy(conn, destConn)
	if err != nil {
		log.Fatalln("unable to copy response to source writer ", err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalln("Unable to listen ", err)
	}
	for {
		conn, err2 := listener.Accept()
		if err2 != nil {
			log.Fatalln("Unable to accept ", err2)
		}
		go handle(conn)
	}
}
