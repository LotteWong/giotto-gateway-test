package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/LotteWong/giotto-gateway-test/dao"
	"github.com/LotteWong/giotto-gateway-test/handler"
	"github.com/LotteWong/giotto-gateway-test/secret"
)

func HttpsServerRun(address string, port int, weight int, id, name string) {
	svc := &dao.Service{
		ID:      id,
		Name:    name,
		Address: address,
		Port:    port,
		Meta:    map[string]string{"weight": fmt.Sprintf("%d", weight)},
	}
	dao.ServiceRegister(svc, false)

	http.HandleFunc("/test_url_rewrite/test", handler.HelloWorldHandler)
	http.HandleFunc("/ping", handler.WrapPingPongHandler(address, port))
	http.HandleFunc("/check", handler.WrapHealthCheckHandler(address, port))

	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf("%s:%d", address, port), secret.Path("server.crt"), secret.Path("server.key"), nil))
}

var address = flag.String("addr", "127.0.0.1", "please input addr")
var port = flag.Int("port", 0, "please input port")
var weight = flag.Int("weight", 0, "please input weight")
var id = flag.String("id", "", "please input id")
var name = flag.String("name", "", "please input name")

func main() {
	flag.Parse()

	HttpsServerRun(*address, *port, *weight, *id, *name)
}
