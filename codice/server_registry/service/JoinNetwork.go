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

	peerAddress := req.Addr
	//Gestiamo, nel caso locale, l'arrivo di un peer crash-recovery e di un nuovo peer
	if *shared.Localhostflag {
		//Verifico se il Peer era già presente nella rete:
		recoveryPeer := -1
		for i := range s.PeerList {
			if peerAddress == s.PeerList[i].Addr {
				recoveryPeer = int(s.PeerList[i].Id)
			}
		}
		if recoveryPeer != -1 {
			return s.sendResponse(int32(recoveryPeer))
		}
	} else if *shared.DockerFlag && req.RecoveryString == "Recovery" {
		log.Printf("Peer with id: %d, is a recovery status", req.RecoveryId)

		//Aggiorno l'address sul peer e invio l'aggiornamento a tutti i peer:
		recoverPeer := &registry.PeerInfo{
			Id:   int32(req.RecoveryId),
			Addr: peerAddress,
		}
		s.recoveryAddress(req)
		UpdatePeer(recoverPeer, s.PeerList, "recoveryPeer")
		log.Println("Peer updated")
		return s.sendResponse(req.RecoveryId)
	}
	//In tutti gli altri casi si tratta di un nuovo peer:
	shared.Id += 1

	// Aggiungi il nuovo Peer alla lista dei Peer registrati
	newPeer := &registry.PeerInfo{
		Id:   int32(shared.Id),
		Addr: peerAddress,
	}
	s.PeerList = append(s.PeerList, newPeer)

	reply, _ := s.sendResponse(int32(shared.Id))

	//Informo tutti i Peer(tranne il nuovo arrivato) del nuovo peer
	UpdatePeer(newPeer, s.PeerList, "newPeer")

	//Aggiungo il nuovo Peer sul file yaml --> In questo modo ho un punto di recupero nel caso in cui il peer registry dovesse crashare
	peerList.AddIntoYaml(newPeer)

	// Attendere il risultato dalla goroutine
	return reply, nil
}

func (s *RegistryServer) recoveryAddress(req *pb.JoinRequest) {
	for i := range s.PeerList {
		if req.RecoveryId == s.PeerList[i].Id {
			s.PeerList[i].Addr = req.Addr
		}
	}
}

func (s *RegistryServer) sendResponse(id int32) (*pb.JoinReply, error) {
	// Prepara la risposta contenente l'ID generato e la lista aggiornata dei peer
	reply := &registry.JoinReply{
		Id:       id,
		PeerList: s.PeerList,
	}
	return reply, nil
}
