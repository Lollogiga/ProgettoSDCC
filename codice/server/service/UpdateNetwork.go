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
	go func() {
		if req.UpdateString == "newPeer" {
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
		} else if req.UpdateString == "recoveryPeer" {
			log.Printf("The peer with id %d is in recovery mode", req.PeerList.GetId())
			log.Printf("Recovery List:\n")
			for i := range shared.PeerList {
				if shared.PeerList[i].Id == req.PeerList.GetId() {
					shared.PeerList[i].Addr = req.PeerList.GetAddr()
				}
				log.Printf("Id: %d Addr: %s, Leader: %t", shared.PeerList[i].Id, shared.PeerList[i].Addr, shared.PeerList[i].Leader)
			}
		}

	}()
	return nil, nil
}
