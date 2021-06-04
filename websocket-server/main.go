package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/LotteWong/giotto-gateway-test/dao"
	"github.com/LotteWong/giotto-gateway-test/handler"
)

func WebSocketServerRun(address string, port int, weight int, id, name string) {
	svc := &dao.Service{
		ID:      id,
		Name:    name,
		Address: address,
		Port:    port,
		Meta:    map[string]string{"weight": fmt.Sprintf("%d", weight)},
	}
	dao.ServiceRegister(svc, false)

	http.HandleFunc("/echo", handler.EchoHandler)
	http.HandleFunc("/home", handler.HomeHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", address, port), nil))
}

var address = flag.String("addr", "127.0.0.1", "please input addr")
var port = flag.Int("port", 0, "please input port")
var weight = flag.Int("weight", 0, "please input weight")
var id = flag.String("id", "", "please input id")
var name = flag.String("name", "", "please input name")

func main() {
	flag.Parse()

	WebSocketServerRun(*address, *port, *weight, *id, *name)
}
