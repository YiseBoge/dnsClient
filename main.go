package main

import (
	"dnsClient/api"
	"dnsClient/models"
	"fmt"
	"log"
)

func main() {

	client := api.ServerClient()

	//domain := models.DomainName{Name:"www.aait.gov.et",Address:"10.90.10.70"}
	//err := domain.Register(client)
	//if err != nil{
	//	log.Fatal(err)
	//}

	//allResults, err := models.DomainName{}.FindAllLocal(client)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("All Local: ", allResults)

	lookupResults, err := models.DomainName{}.Lookup(client, "www.aait.gov.et")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Looked up: ", lookupResults)

	domain2 := models.DomainName{Name: "www.aait.gov.et", Address: "10.90.10.70"}
	err2 := domain2.Remove(client)
	if err2 != nil {
		log.Fatal(err)
	}

	//var result bool
	//domain := models.DomainName{Name:"www.apple.com",Address:"1.2.3.4"}
	//err = client.Call("API.Remove", domain, &result)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//fmt.Println("Removed: ", result)

	//var result bool
	//domain := models.DomainName{Name: "www.asp.net", Address: "34.24.14.4"}
	//err = client.Call("API.RemoveCache", domain, &result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Removed: ", result)
	//
	//err = client.Call("API.Lookup", "www.asp.net", &results)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Looked up: ", results)
}
