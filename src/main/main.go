package main

import (
	"log"
	"net"
	"os"
)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {

	log.Print("Starting program...")

	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)

	if err != nil {
		log.Print("Error while serving connections")
	}

	defer listener.Close()
	log.Print("Listening on ", CONN_HOST, " : ", CONN_PORT)

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Print("Error accepting : ", err.Error())
			os.Exit(1)
		}

		go serverClientConnection(connection)

	}

}

func serverClientConnection(conn net.Conn) {

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
