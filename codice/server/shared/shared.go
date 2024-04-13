package shared

import "codice/server/registry"

// Variabili di configurazione:

var PeerList []registry.PeerInfo
var Port int
var MyId, LeaderId int32

// variabili per algoritmo di Dolev-Klawe-Rodeh
var Token, MyToken = 0, -1
var NumNode = 0
