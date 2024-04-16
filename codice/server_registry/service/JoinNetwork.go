package service

import (
	"codice/server_registry/peerList"
	"codice/server_registry/registry"
	pb "codice/server_registry/registry"
	"codice/server_registry/shared"
	"context"
	"log"
)

type RegistryServer struct {
	pb.UnimplementedRegistryServer
	PeerList []*pb.PeerInfo
}

func (s *RegistryServer) JoinNetwork(ctx context.Context, req *registry.JoinRequest) (*registry.JoinReply, error) {

	//Ottengo dal messaggio Join Request l'indirizzo del client
	address := req.Addr
	log.Printf("New peer join the network: %s\n", address)

	//Verifico se il Peer era già presente nella rete:
	recoveryPeer := -1
	for i := range s.PeerList {
		if address == s.PeerList[i].Addr {
			recoveryPeer = int(s.PeerList[i].Id)
		}
	}
	if recoveryPeer != -1 {
		// Prepara la risposta contenente l'ID generato e la lista aggiornata dei peer
		reply := &registry.JoinReply{
			Id:       int32(recoveryPeer),
			PeerList: s.PeerList,
		}
		return reply, nil
	}

	//Altrimenti è un nuovo peer:
	shared.Id += 1
	// Aggiungi il nuovo Peer alla lista dei Peer registrati
	newPeer := &registry.PeerInfo{
		Id:   int32(shared.Id),
		Addr: address,
	}
	s.PeerList = append(s.PeerList, newPeer)

	// Prepara la risposta contenente l'ID generato e la lista aggiornata dei peer
	reply := &registry.JoinReply{
		Id:       int32(shared.Id),
		PeerList: s.PeerList,
	}

	//Informo tutti i Peer(tranne il nuovo arrivato) del nuovo peer
	UpdatePeer(newPeer, s.PeerList)

	//Aggiungo il nuovo Peer sul file yaml --> In questo modo ho un punto di recupero nel caso in cui il server registry dovesse crashare
	peerList.AddIntoYaml(newPeer)

	// Attendere il risultato dalla goroutine
	return reply, nil
}
