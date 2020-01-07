package main

import (
	"dnsClient/api"
	"dnsClient/models"
	"fmt"
	"log"
)

func main() {

	var results []models.DomainName

	client := api.ServerClient()

	err := client.Call("API.FindAll", "", &results)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Database: ", results)

	err = client.Call("API.Lookup", "www.apple.com", &results)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Looked up: ", results)
}
