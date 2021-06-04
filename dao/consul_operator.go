package dao

import (
	"fmt"
	"log"

	consulapi "github.com/hashicorp/consul/api"
)

type Service struct {
	ID      string
	Name    string
	Address string
	Port    int
	Meta    map[string]string
}

func ServiceRegister(svc *Service, doHealthCheck bool) {
	// connect consul client
	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalf("consul client error: %v", err)
	}

	// register service info
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = svc.ID
	registration.Name = svc.Name
	registration.Address = svc.Address
	registration.Port = svc.Port
	registration.Meta = svc.Meta

	// register health check
	if doHealthCheck {
		registration.Check = &consulapi.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/check"),
			Timeout:                        "3s",
			Interval:                       "5s",
			DeregisterCriticalServiceAfter: "30s",
		}
	}

	// register to consul
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("register server error: %v", err)
	}
}
