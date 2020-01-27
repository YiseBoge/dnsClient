package api

import (
	"dnsClient/config"
	"net"
	"net/rpc"
	"time"
)

func ServerClient() (*rpc.Client, error) {
	address := config.LoadConfig().Server.Address
	port := config.LoadConfig().Server.Port

	timeout := 5 * time.Second
	_, err := net.DialTimeout("tcp", address+":"+port, timeout)
	if err != nil {
		return nil, err
	}

	client, err := rpc.DialHTTP("tcp", address+":"+port)
	if err != nil {
		return nil, err
	}
	return client, nil
}
