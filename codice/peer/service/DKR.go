package service

import (
	pb "codice/peer/registry"
	"codice/peer/shared"
	"google.golang.org/grpc"
	"log"
	"math"
)

import (
	"context"
)

func (s *Election) DKRElection(ctx context.Context, req *pb.ElectionRequest) (*pb.ElectionReply, error) {
	if req.Election == "Election" {
		//Ho ricevuto dal mio predecessore un stimateLeader, lo confronto con il mio:
		//Se l'id ricevuto, è maggiore del mio stimateLeader, aggiorno il mio EstimateLeader:
		if shared.State == "waiting" && req.ElectionId == shared.MyId {
			BecomeleaderDKR()
			return nil, nil
		}

		if req.ElectionId > shared.MyId {
			shared.StimateLeader = req.ElectionId
			shared.State = "passive"
		} else {
			shared.State = "waiting"
		}
		//Propago l'informazione
		DKR()
		return nil, nil
	}
	//Altrimenti ho un messaggio di Update:
	//Ripristino stato e StimateLeader:
	shared.StimateLeader = shared.MyId
	shared.State = "active"
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
func DKR() {
	//Mi calcolo il mio successore nella rete:
	nextNode := math.Mod(float64(shared.MyId+1), float64(shared.NumNode))

	for {
		//Provo a connettermi al mio successore:
		conn, err := grpc.Dial(shared.PeerList[int32(nextNode)].Addr, grpc.WithInsecure())
		defer conn.Close()
		//Se la connessione non va a buon fine, invio al successore del mio successore.
		if err != nil {
			log.Printf("Peer Unreachble")
			nextNode = math.Mod(nextNode+1, float64(shared.NumNode))
		}

		//Provo ad inviare la mia variabile EstimateLeader al successore:
		client := pb.NewElectionClient(conn)
		log.Printf("Sto per inviare la mia variabile leader: %d", shared.StimateLeader)
		_, err = client.DKRElection(context.Background(), &pb.ElectionRequest{
			Election:   "Election",
			ElectionId: shared.StimateLeader,
		})
		//Se l'invio non va a buon fine provo ad inviare al successore del mio successore:
		if err != nil {
			log.Printf("Errore durante l'invio della variabile al nodo\n")
			nextNode = math.Mod(nextNode+1, float64(shared.NumNode))
		} else {
			break
		}
	}
	return
}

func BecomeleaderDKR() {
	//Setto tutti i leader a false
	shared.State = "active"
	shared.StimateLeader = shared.MyId
	for i := range shared.PeerList {
		if shared.PeerList[i].Id == shared.MyId {
			shared.PeerList[i].Leader = true
		} else {
			shared.PeerList[i].Leader = false
		}
	}
	log.Printf("I'm new Leader")
	//Devo avvisare tutti i peer che sono il nuovo Leader:
	for _, server := range shared.PeerList {
		conn, err := grpc.Dial(server.Addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to Peer: %v", err)
		}
		defer conn.Close()

		client := pb.NewElectionClient(conn)

		//Invio la nuova entry al client
		_, err = client.DKRElection(context.Background(), &pb.ElectionRequest{
			Election:   "Update",
			ElectionId: shared.MyId,
		})
	}
}
