package main

import (
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func main() {

	http.Handle("/", websocket.Handler(serverClientConnection))
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Print("Error while serving connections")
	}

}

func serverClientConnection(ws *websocket.Conn) {

	addr := ws.RemoteAddr().String()
	log.Print(addr)

	bytesToRead := make([]byte, 1024)
	log.Print(ws.Request().Header)

	log.Print("Headers are : ")
	for k, v := range ws.Request().Header {
		log.Print(k, " |======|  ", v)
	}

	for {
		read, err := ws.Read(bytesToRead)
		if err != nil {
			log.Print("Error")
			log.Print(err)
			break
		}
		log.Print("read bytes : ", read)
	}

}
