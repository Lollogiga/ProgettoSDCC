package service

import (
	pb "codice/server/registry"
	"codice/server/shared"
	"context"
	"log"
)

type UpdateServer struct {
	pb.UnimplementedUpdateServer
}

// Aggiornamento lista a seguito di un nuovo peer:
func (s *UpdateServer) UpdateNetwork(ctx context.Context, req *pb.UpdateMessage) (*pb.UpdateResponse, error) {

	log.Printf("New Server join Network: %s\n", req)
	shared.NumNode++
	shared.PeerList = append(shared.PeerList, pb.PeerInfo{
		Id:     req.PeerList.GetId(),
		Addr:   req.PeerList.GetAddr(),
		Leader: false,
	})

	log.Printf("Update List:\n")
	for _, server := range shared.PeerList {
		log.Printf("Id: %d Addr: %s, Leader: %t", server.Id, server.Addr, server.Leader)
	}

	return nil, nil
}
