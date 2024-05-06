package service

import (
	"codice/server_registry/registry"
	"context"
	"google.golang.org/grpc"
	"log"
)

// Aggiorno tutti i peer del nuovo arrivato
func UpdatePeer(newPeer *registry.PeerInfo, PeerList []*registry.PeerInfo, updateString string) {
	log.Printf("The id in Update Peer is: %d", newPeer.Id)

	lenPeerList := len(PeerList)
	if updateString == "newPeer" {
		lenPeerList = len(PeerList) - 1
	}

	for i := 0; i < lenPeerList; i++ {
		log.Printf("Connect to peer with address: %s", PeerList[i].Addr)
		conn, err := grpc.Dial(PeerList[i].Addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to Peer: %v", err)
		}
		defer conn.Close()

		client := registry.NewUpdateClient(conn)

		//Invio la nuova entry al client
		_, err = client.UpdateNetwork(context.Background(), &registry.UpdateMessage{
			UpdateString: updateString,
			PeerList:     newPeer,
		})
		if err != nil {
			log.Printf("Failed to UpdateNetwork:\nPeer with address: %s unreachable\n", PeerList[i].Addr)

		}

	}
	log.Printf("All peers have been updated\n")
}
