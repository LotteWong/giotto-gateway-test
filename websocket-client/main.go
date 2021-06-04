package main

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

var jwt = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjI2OTAyNjIsImlzcyI6ImFwcF9pZF9hIn0.Sdzkjm_N2nLONzEHCCQleDXk5XBMpVR7PJSEUN944Aw"
var url = "ws://127.0.0.1:80/test_websocket_service/echo"
var origin = "http://127.0.0.1:80/test_websocket_service/home"

func WebSocketClientRun(url, jwt, origin string) {
	ws, err := websocket.Dial(url, jwt, origin)
	if err != nil {
		log.Fatal(err)
	}
	message := []byte("Hello, world!")
	_, err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", message)

	var msg = make([]byte, 512)
	m, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", msg[:m])

	ws.Close()
}

func main() {
	WebSocketClientRun(url, jwt, origin)
}
