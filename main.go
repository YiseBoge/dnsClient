package main

import (
	"bufio"
	"dnsClient/api"
	"dnsClient/models"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

func main() {

	client, err := api.ServerClient()
	if err != nil {
		log.Println("Site unreachable, error: ", err)
	}

	//domain := models.DomainName{Name:"www.aait.gov.et",Address:"10.90.10.70"}
	//err := domain.Register(client)
	//if err != nil{
	//	log.Fatal(err)
	//}

	//domain := models.DomainName{Name:"www.eca.et",Address:"10.20.30.40"}
	//err := domain.Register(client)
	//if err != nil{
	//	log.Fatal(err)
	//}

	//allResults, err := models.DomainName{}.FindAllLocal(client)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("All Local: ", allResults)

	//lookupResults, err := models.DomainName{}.Lookup(client, "www.eca.et")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Looked up: ", lookupResults)

	//domain2 := models.DomainName{Name: "www.aait.gov.et", Address: "10.90.10.70"}
	//err2 := domain2.Remove(client)
	//if err2 != nil {
	//	log.Fatal(err2)
	//}
	fmt.Println("Welcome to DomInator Client.")
	for true {
		fmt.Printf(">> ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		list := strings.Split(scanner.Text(), " ")

		if len(list) < 2 {
			fmt.Println("Usage, `command <args>`")
		} else {
			Respond(client, list[0], list[1:])
		}
	}
}

func Respond(client *rpc.Client, command string, args []string) {

	switch command {
	case "lookup":
		if len(args) != 1 {
			fmt.Println("Usage, `lookup <name>`")
			break
		}
		lookupResults, err := models.DomainName{}.Lookup(client, args[0])
		if err != nil {
			log.Println(err)
		}
		var addresses []string
		for _, result := range lookupResults {
			addresses = append(addresses, result.Address)
		}
		fmt.Printf("IP Addresses for `%s` are %s\n", args[0], addresses)
	case "register":
		if len(args) != 2 {
			fmt.Println("Usage, `register <name> <address>`")
			break
		}
		domain := models.DomainName{Name: args[0], Address: args[1]}
		err := domain.Register(client)
		if err != nil {
			log.Println(err)
			fmt.Println("Usage, `register <name> <address>`")
			break
		}
		fmt.Println("Successfully Registered!!")
	case "remove":
		if len(args) != 2 {
			fmt.Println("Usage, `remove <name> <address>`")
			break
		}
		domain := models.DomainName{Name: args[0], Address: args[1]}
		err := domain.Remove(client)
		if err != nil {
			log.Println(err)
			fmt.Println("Usage, `remove <name> <address>`")
			break
		}
		fmt.Println("Successfully Removed!!")
	default:
		fmt.Println("Usage, `command <args>`")
		fmt.Println("Allowed commands are: [`lookup`, `register`, `remove`]")
	}
}
