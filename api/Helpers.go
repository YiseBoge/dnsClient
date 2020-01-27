package api

import (
	"dnsClient/config"
	"log"
	"net"
	"net/rpc"
	"time"
)

func ServerClient() *rpc.Client {
	address := config.LoadConfig().Server.Address
	port := config.LoadConfig().Server.Port

	timeout := 5 * time.Second
	_, err := net.DialTimeout("tcp", address+":"+port, timeout)
	if err != nil {
		log.Fatal("Site unreachable, error: ", err)
	}

	client, err := rpc.DialHTTP("tcp", address+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
