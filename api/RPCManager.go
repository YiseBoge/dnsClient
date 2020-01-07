package api

import (
	"dnsClient/config"
	"log"
	"net/rpc"
)

func ServerClient() *rpc.Client {
	parentAddress := config.LoadConfig().Server.Address
	parentPort := config.LoadConfig().Server.Port

	client, err := rpc.DialHTTP("tcp", parentAddress+":"+parentPort)
	if err != nil{
		log.Fatal(err)
	}
	return client
}
