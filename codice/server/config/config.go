package config

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

var ServerAddress string
var MyAddress string
var BullySelected bool
var DolevSelected bool

// Struttura per memorizzare le informazioni del service registry
type ServiceRegistry struct {
	Address string `json:"address"`
	Port    string `json:"port"`
}

// Struttura per memorizzare le informazioni dell'algoritmo
type Algorithm struct {
	Bully bool `json:"Bully"`
	Dolev bool `json:"Dolev"`
}

// Struttura per memorizzare le informazioni di ogni tipo di configurazione (localhost e docker)
type Config struct {
	Localhost LocalhostConfig `json:"localhost"`
	Docker    DockerConfig    `json:"docker"`
}

// Struttura per memorizzare le informazioni di localhost
type LocalhostConfig struct {
	ServiceRegistry ServiceRegistry `json:"service_registry"`
	Algorithm       Algorithm       `json:"algorithm"`
}

// Struttura per memorizzare le informazioni di docker
type DockerConfig struct {
	ServiceRegistry ServiceRegistry `json:"service_registry"`
	Peer            ServiceRegistry `json:"peer"`
	Algorithm       Algorithm       `json:"algorithm"`
}

func DockerConfiguration() {
	fileContent, err := os.ReadFile("../config.json")

	if err != nil {
		fmt.Println("Errore nella lettura del file:", err)
		return
	}

	var configData Config
	err = json.Unmarshal(fileContent, &configData)
	if err != nil {
		fmt.Println("Errore nel parsing del file JSON:", err)
		return
	}
	ServerAddress = configData.Docker.ServiceRegistry.Address + configData.Docker.ServiceRegistry.Port
	//MyAddress = configData.Docker.Peer.Address + configData.Docker.Peer.Port
	BullySelected = configData.Docker.Algorithm.Bully
	DolevSelected = configData.Docker.Algorithm.Dolev

	hostname := "peer" // Sostituisci con il nome host desiderato

	// Risolvi il nome host in un indirizzo IP
	addrs, err := net.LookupHost(hostname)
	if err != nil {
		fmt.Printf("Errore durante la risoluzione dell'indirizzo IP per %s: %v\n", hostname, err)
		return
	}

	// Utilizza il primo indirizzo IP risolto
	ip := addrs[0]

	// Stampa l'indirizzo IP
	fmt.Printf("Indirizzo IP per %s: %s\n", hostname, ip)
	MyAddress = ip + configData.Docker.Peer.Port
}

func LocalConfig() {

	fileContent, err := os.ReadFile("../config.json")

	if err != nil {
		fmt.Println("Errore nella lettura del file:", err)
		return
	}

	var configData Config
	err = json.Unmarshal(fileContent, &configData)
	if err != nil {
		fmt.Println("Errore nel parsing del file JSON:", err)
		return
	}
	ServerAddress = configData.Localhost.ServiceRegistry.Address + configData.Localhost.ServiceRegistry.Port
	log.Printf("%s", ServerAddress)
	BullySelected = configData.Localhost.Algorithm.Bully
	DolevSelected = configData.Localhost.Algorithm.Dolev
}
