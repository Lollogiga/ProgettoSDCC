package service

import (
	pb "codice/peer/registry"
	"codice/peer/shared"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

type Election struct {
	pb.UnimplementedElectionServer
}

func (s *Election) BullyElection(ctx context.Context, req *pb.ElectionRequest) (*pb.ElectionReply, error) {

	log.Printf("Message: %s", req.Election)
	if req.Election == "Election" {

		//Ricevo l'indirizzo del peer che richiede l'elezione
		messageOk := fmt.Sprint("OK")
		reply := &pb.ElectionReply{
			ElectionReply: messageOk,
		}
		Bully()

		return reply, nil
	}

	//Altrimenti è un messaggio di coordinator:
	for i := range shared.PeerList {
		if shared.PeerList[i].Id == req.ElectionId {
			shared.PeerList[i].Leader = true
			shared.LeaderId = shared.PeerList[i].Id
		} else {
			shared.PeerList[i].Leader = false
		}
		log.Printf("After Update message peer: %d have leader: %t:", shared.PeerList[i].Id, shared.PeerList[i].Leader)
	}
	return nil, nil
}

func Bully() {
	log.Printf("Avvio elezione Bully\n")
	for _, server := range shared.PeerList {
		if server.Id > shared.MyId {
			//Invio messaggio Election:

			conn, err := grpc.Dial(server.Addr, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("Failed to connect to Peer: %v", err)
			}
			defer conn.Close()

			client := pb.NewElectionClient(conn)

			reply, err := client.BullyElection(context.Background(), &pb.ElectionRequest{
				Election:   "Election",
				ElectionId: shared.MyId,
			})
			if err != nil {
				continue
			}
			if reply.ElectionReply == "OK" {
				log.Printf("I can't became leader\n")
				return
			}
		}
	}

	//Se esco dal for vuol dire che ho vinto le elezioni
	for i := range shared.PeerList {
		if shared.PeerList[i].Id == shared.MyId {
			shared.PeerList[i].Leader = true
		} else {
			shared.PeerList[i].Leader = false
		}
	}
	log.Printf("I'm new leader")

	//Devo avvisare tutti i peer che sono il nuovo leader:
	for _, server := range shared.PeerList {
		conn, err := grpc.Dial(server.Addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to Peer: %v", err)
		}
		defer conn.Close()

		client := pb.NewElectionClient(conn)

		//Invio la nuova entry al client
		_, err = client.BullyElection(context.Background(), &pb.ElectionRequest{
			Election:   "Coordinator",
			ElectionId: shared.MyId,
		})
	}
}
