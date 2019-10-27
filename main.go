package main

import (
	"log"
	"net"
	"strconv"
)

//Will be using these port for port scanning later
/*
var ports = []string{
	"80",
	"21",
	"22",
	"443",
	"8080",
}
*/

func main() {
	ActiveThread := 0
	doneChann := make(chan bool)
	subNet := "192.168.100"

	for i := 0; i < 256; i++ {

		realSubnet := subNet + "." + strconv.Itoa(i)
		go resolvName(realSubnet, doneChann)
		ActiveThread++
	}

	for ActiveThread > 0 {
		<-doneChann
		ActiveThread--
	}
}
func resolvName(addr string, doneChann chan bool) {

	hostname, err := net.LookupAddr(addr)
	if err == nil {
		log.Printf("%v -- %v", addr, hostname)
	}

	doneChann <- true
}
