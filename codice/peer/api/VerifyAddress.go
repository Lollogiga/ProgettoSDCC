package api

import (
	"log"
	"net"
	"strings"
)

// verifico correttezza indirizzo passato in input:
func VerifyAddress(address string) {
	// Controllo dell'indirizzo IP e della porta
	ipPort := strings.Split(address, ":")
	if len(ipPort) != 2 {
		log.Fatalf("Invalid address format. Use IP:port")
	}

	//Verifico che il formato sia corretto:
	ip := net.ParseIP(ipPort[0])
	if ip == nil {
		log.Fatalf("Invalid IP address")
	}

	//Verifico che la porta sia disponibile
	port := ipPort[1]
	if _, err := net.LookupPort("tcp", port); err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	log.Printf("Indirizzo: %s:%s", ip, port)
}
