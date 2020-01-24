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

	//var result bool
	//domain := models.DomainName{Name:"www.apple.com",Address:"1.2.3.4"}
	//err := client.Call("API.Register", domain, &result)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//fmt.Println("Added: ", result)

	err := client.Call("API.FindAll", "", &results)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database: ", results)

	//var result bool
	//domain := models.DomainName{Name:"www.apple.com",Address:"1.2.3.4"}
	//err = client.Call("API.Remove", domain, &result)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//fmt.Println("Removed: ", result)

	var result bool
	domain := models.DomainName{Name: "www.asp.net", Address: "34.24.14.4"}
	err = client.Call("API.RemoveCache", domain, &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Removed: ", result)

	err = client.Call("API.Lookup", "www.asp.net", &results)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Looked up: ", results)
}
