package main

import (
	"fmt"
	"net"
)

func domainLookup(name string) { //Defining domain lookup function
	fmt.Println(ColorYellow, "\n[+] CNAME ")
	cname, _ := net.LookupCNAME(name)
	fmt.Println(ColorGreen, cname) //printing cname
	fmt.Println(ColorYellow, "\n[+] TXT records ")
	txtrecords, _ := net.LookupTXT(name)

	for _, txt := range txtrecords {
		fmt.Println(ColorGreen, txt) //printing txt records
	}
	fmt.Println(ColorYellow, "\n[+] A and AAAA  ")

	iprecords, _ := net.LookupIP(name)
	for _, ip := range iprecords {
		fmt.Println(ColorGreen, ip) //printing AAAA / A records
	}
	fmt.Println(ColorYellow, "\n[+] NS  ")
	nameserver, _ := net.LookupNS(name)
	for _, ns := range nameserver {
		fmt.Println(ColorGreen, ns) //Printing nameservers
	}
	fmt.Println(ColorYellow, "\n[+] MX \n")

	mxrecords, _ := net.LookupMX(name)
	for _, mx := range mxrecords {
		fmt.Println(ColorGreen, mx.Host, mx.Pref) // Printing MX records
	}
}
