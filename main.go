package main

import (
	"bufio"
	"dnsClient/api"
	"dnsClient/config"
	"dnsClient/models"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"regexp"
	"strings"
)

func main() {

	fmt.Println("Welcome to the DomaInator Client.")
	configuration := config.LoadConfig()

	portRegex, _ := regexp.Compile("^([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])(?::([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5]))?$")
	domainRegex, _ := regexp.Compile("^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9].[a-zA-Z]{2,}$")
	ipRegex, _ := regexp.Compile("^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]).){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$")

	var res3 string
	for true {
		fmt.Printf("DNS Server address = \"%s\" press 'Enter' to continue or provide new address: ", configuration.Server.Address)
		_, _ = fmt.Scanln(&res3)

		if res3 == "" {
			break
		}

		if domainRegex.MatchString(res3) || ipRegex.MatchString(res3) {
			configuration.Server.Address = res3
			break
		}
		fmt.Println("**Bad input, Please try again**")
	}

	var res4 string
	for true {
		fmt.Printf("DNS Server port = \"%s\" press 'Enter' to continue or provide new port: ", configuration.Server.Port)
		_, _ = fmt.Scanln(&res4)

		if res4 == "" {
			break
		}

		if portRegex.MatchString(res4) {
			configuration.Server.Port = res4
			break
		}
		fmt.Println("**Bad input, Please try again**")
	}

	config.SaveConfig(configuration)

	client, err := api.ServerClient()
	if err != nil {
		log.Fatal("Site unreachable, error: ", err)
	}

	for true {
		fmt.Printf(">> ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		str := strings.TrimSpace(scanner.Text())
		if str == "exit" || str == "stop" {
			break
		}
		list := strings.Split(str, " ")

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
			break
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
