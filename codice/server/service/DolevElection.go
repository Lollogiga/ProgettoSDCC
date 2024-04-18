package service

import (
	pb "codice/server/registry"
	"codice/server/shared"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math"
)

func (s *Election) DolevElection(ctx context.Context, req *pb.ElectionRequest) (*pb.ElectionReply, error) {
	//Quando ricevo un messaggio di elezione, può essere un messaggio di Token o di Election:
	log.Printf("Dolev Message: %s", req.Election)
	if req.Election == "Token" {
		//Ricevo il token del peer che mi precede nel ring
		messageOk := fmt.Sprint("OK")
		reply := &pb.ElectionReply{
			ElectionReply: messageOk,
		}
		//Se il token che ho ricevuto è minore del mio token, allora non propago l'elezione (ne ho già avviata una io o un peer con id maggiore)
		if req.TokenId < int32(shared.MyToken) {
			return reply, nil
		}
		shared.MyToken = int(req.TokenId)
		Dolev()
		return reply, nil
	}

	//Altrimenti è un messaggio di update:
	shared.MyToken = -1
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

// In tal caso io sto avviando una nuova elezione, mi genero un mio token.
func DolevStartElection() {
	log.Printf("Avvio Elezione Dolev\n")
	shared.MyToken = int(shared.MyId)
	Dolev()
}

func Dolev() {
	//Mi calcolo il mio successore nel ring:
	nextNode := math.Mod(float64(shared.MyId+1), float64(shared.NumNode))
	//Verifico se il mio id è quello maggiore nella rete:
	for _, server := range shared.PeerList {
		if server.Id > shared.MyId {
			//Se esiste un nodo che ha id maggiore del mio, invio il token al mio successivo nella rete:
			for {
				conn, err := grpc.Dial(shared.PeerList[int32(nextNode)].Addr, grpc.WithInsecure())
				log.Printf("After Dial\n")
				defer conn.Close()
				if err != nil {
					log.Printf("Peer Unreachble")
					nextNode = math.Mod(nextNode+1, float64(shared.NumNode))
					//Devo passare al mio successore se e solo se quest'ultimo ha id strettamente maggiore del mio:
					if shared.PeerList[int32(nextNode)].Id <= shared.MyId {
						Becomeleader()
						return
					} else {
						continue
					}
				}

				//Se è riuscita la connessione provo a inviare il token al mio successore:
				client := pb.NewElectionClient(conn)
				//Invio Token al nodo successivo:
				_, err = client.DolevElection(context.Background(), &pb.ElectionRequest{
					Election: "Token",
					//ElectionId: shared.MyId,
					TokenId: shared.MyId,
				})
				if err != nil {
					log.Printf("Errore durante l'invio del token al nodo\n")
					nextNode = math.Mod(nextNode+1, float64(shared.NumNode))
					//Devo passare al mio successore se e solo se quest'ultimo ha id strettamente maggiore del mio:
					if shared.PeerList[int32(nextNode)].Id <= shared.MyId {
						Becomeleader()
						return
					} else {
						continue
					}
				} else {
					break
				}
			}
			//Ho inviato il token al mio successore:
			return
		}
	}
	Becomeleader()
	return
}

func Becomeleader() {
	//Setto tutti i leader a false
	shared.MyToken = -1
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
		_, err = client.DolevElection(context.Background(), &pb.ElectionRequest{
			Election:   "Update",
			ElectionId: shared.MyId,
		})
	}
}
