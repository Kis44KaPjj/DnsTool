package main

import (
	"fmt"
	"net"
)

var version = "dev"

func main() {
	showbanner()

	fmt.Println("[1] - Resolve Domain")
	fmt.Println("[2] - Reverse Lookup")
	fmt.Println("[3] - NameServers")
	fmt.Println("[0] - Exit")

	for {
		var choice string
		fmt.Print("Select Option: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			ResolveDomain()
		case "2":
			ReverseLookup()
		case "3":
			LookupNameServers()
		case "0":
			fmt.Println("Exiting...!")
			return
		default:
			fmt.Println("Invalid Option:", choice)
		}
	}
}

func ResolveDomain() {
	var InputDomain string
	fmt.Print("\nEnter a domain: ")
	fmt.Scanln(&InputDomain)

	ips, err := net.LookupIP(InputDomain)
	if err != nil {
		fmt.Println("Error resolving domain", err)
		return
	}

	for _, ip := range ips {
		fmt.Printf("IP for %s => %s\n", InputDomain, ip)
	}
	fmt.Println()
}

func ReverseLookup() {
	var InputIP string
	fmt.Print("\nEnter IP: ")
	fmt.Scanln(&InputIP)

	domains, err := net.LookupAddr(InputIP)
	if err != nil {
		fmt.Println("Reverse lookup failed:", err)
		return
	}

	for _, name := range domains {
		fmt.Printf("Domain for %s => %s\n", InputIP, name)
	}
	fmt.Println()

}

func LookupNameServers() {
	var InputDomain string
	fmt.Print("\nEnter a domain: ")
	fmt.Scanln(&InputDomain)

	nsrecords, err := net.LookupNS(InputDomain)
	if err != nil {
		fmt.Println("Error finding NS records:", err)
		return
	}

	for _, ns := range nsrecords {
		fmt.Printf("NameServers for %s => %s\n", InputDomain, ns.Host)
	}
	fmt.Println()
}
