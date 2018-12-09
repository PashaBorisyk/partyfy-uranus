package main

import (
	"log"
	"net"
	"os"
)

const (
	ConnHost = "0.0.0.0"
	ConnPort = "3333"
	ConnType = "tcp"
)

func main() {

	log.Print("Starting program...")

	listener, err := net.Listen(ConnType, ConnHost+":"+ConnPort)

	if err != nil {
		log.Print("Error while serving connections")
	}

	defer listener.Close()
	log.Print("Listening on ", ConnHost, " : ", ConnPort)

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Print("Error accepting : ", err.Error())
			os.Exit(1)
		}

		go serveClientInputConnection(connection)
		go serveClientOutputConnection(connection)

	}

}

func serveClientInputConnection(conn net.Conn) {

	log.Print("Connection from ", conn.RemoteAddr())

	buf := make([]byte, 1024)

	for {
		read, err := conn.Read(buf)
		if err != nil {
			log.Print("Error")
			log.Print(err)
			break
		}
		log.Print("read bytes : ", read)
	}

}

func serveClientOutputConnection(conn net.Conn) {

}
