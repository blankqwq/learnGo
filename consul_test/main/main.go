package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/hashicorp/consul/api"
)

func Registe(name string, id string, address string, port int, tags []string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", "localhost", 9001)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	register := api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Tags:    tags,
		Port:    port,
		Address: address,
		Check: &api.AgentServiceCheck{
			Name:                           name,
			Interval:                       "5s",
			Timeout:                        "5s",
			HTTP:                           fmt.Sprintf("http://%s:%d/api/health", address, port),
			DeregisterCriticalServiceAfter: "5s",
		},
	}

	err = client.Agent().ServiceRegister(&register)
	if err != nil {
		panic(err)
	}
	return nil
}

func main() {
	err := Registe("user-web", "user-web", "192.168.233.119", 9003, []string{"user", "golang", "web"})
	if err != nil {
		println(err.Error())
		return
	}
}
