package peerList

import (
	pb "codice/server_registry/registry"
	"codice/server_registry/shared"
	"fmt"
)

// Aggiungo un nuovo peer sul file Yaml seguendo la sintassi yaml:
func AddIntoYaml(newPeer *pb.PeerInfo) {
	newPeerEntry := fmt.Sprintf("  - id: %d\n"+
		"    addr: %s", newPeer.Id, newPeer.Addr)

	_, err := fmt.Fprintln(shared.File, newPeerEntry)
	if err != nil {
		panic(err)
	}
}
