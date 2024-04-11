package main

import (
	pb "codice/server_registry/registry"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type registryServer struct {
	pb.UnimplementedRegistryServer
	PeerList []*pb.PeerInfo
}

var file *os.File
var id = -1
var PeerList []pb.PeerInfo

// Gestiso l'ingresso di nuovi peer con una goRoutine
/*func (s *registryServer) JoinNetwork(ctx context.Context, req *pb.JoinRequest) (*pb.JoinReply, error) {
	// Crea un canale per comunicare il risultato della goroutine
	resultChan := make(chan *pb.JoinReply)
	defer close(resultChan) // Chiudi il canale quando la funzione termina

	go func() {
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
			reply := &pb.JoinReply{
				Id:       int32(recoveryPeer),
				PeerList: s.PeerList,
			}
			resultChan <- reply // Invia la risposta al canale
			return
		}
		//Altrimenti è un nuovo peer:
		id += 1
		// Aggiungi il nuovo Peer alla lista dei Peer registrati
		newPeer := &pb.PeerInfo{
			Id:   int32(id),
			Addr: address,
		}
		s.PeerList = append(s.PeerList, newPeer)

		// Prepara la risposta contenente l'ID generato e la lista aggiornata dei peer
		reply := &pb.JoinReply{
			Id:       int32(id),
			PeerList: s.PeerList,
		}

		//Informo tutti i Peer(tranne il nuovo arrivato) del nuovo peer
		UpdatePeer(newPeer, s.PeerList)

		//Aggiungo il nuovo Peer sul file yaml --> In questo modo ho un punto di recupero nel caso in cui il server registry dovesse crashare
		s.addIntoYaml(newPeer)

		resultChan <- reply // Invia la risposta al canale
	}()

	// Attendere il risultato dalla goroutine
	reply := <-resultChan
	return reply, nil
}*/

func (s *registryServer) JoinNetwork(ctx context.Context, req *pb.JoinRequest) (*pb.JoinReply, error) {

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
		reply := &pb.JoinReply{
			Id:       int32(recoveryPeer),
			PeerList: s.PeerList,
		}
		return reply, nil
	}

	//Altrimenti è un nuovo peer:
	id += 1
	// Aggiungi il nuovo Peer alla lista dei Peer registrati
	newPeer := &pb.PeerInfo{
		Id:   int32(id),
		Addr: address,
	}
	s.PeerList = append(s.PeerList, newPeer)

	// Prepara la risposta contenente l'ID generato e la lista aggiornata dei peer
	reply := &pb.JoinReply{
		Id:       int32(id),
		PeerList: s.PeerList,
	}

	//Informo tutti i Peer(tranne il nuovo arrivato) del nuovo peer
	UpdatePeer(newPeer, s.PeerList)

	//Aggiungo il nuovo Peer sul file yaml --> In questo modo ho un punto di recupero nel caso in cui il server registry dovesse crashare
	s.addIntoYaml(newPeer)

	// Attendere il risultato dalla goroutine
	return reply, nil
}

// Aggiungo un nuovo peer sul file Yaml seguendo la sintassi yaml:
func (s *registryServer) addIntoYaml(newPeer *pb.PeerInfo) {
	newPeerEntry := fmt.Sprintf("  - id: %d\n"+
		"    addr: %s", newPeer.Id, newPeer.Addr)

	_, err := fmt.Fprintln(file, newPeerEntry)
	if err != nil {
		panic(err)
	}
}

// Aggiorno tutti i peer del nuovo arrivato
func UpdatePeer(newPeer *pb.PeerInfo, PeerList []*pb.PeerInfo) {
	length := len(PeerList) - 1
	for i := 0; i < length; i++ {
		log.Printf("Connect to peer with address: %s", PeerList[i].Addr)
		conn, err := grpc.Dial(PeerList[i].Addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to Peer: %v", err)
		}
		defer conn.Close()

		client := pb.NewUpdateClient(conn)

		//Invio la nuova entry al client
		_, err = client.UpdateNetwork(context.Background(), &pb.UpdateMessage{
			PeerList: newPeer,
		})
		if err != nil {
			log.Printf("Failed to UpdateNetwork:\nPeer with address: %s unreachable\n", PeerList[i].Addr)

		}

	}
	log.Printf("All peers have been updated\n")
}

func main() {

	server := &registryServer{
		PeerList: []*pb.PeerInfo{},
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRegistryServer(s, server)
	log.Println("Registry listening on port 50051")

	file, err = os.OpenFile("ServerList.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	//file, err = os.OpenFile("ServerList.yaml", os.O_RDWR, 0600)

	if err != nil {
		log.Fatalf("error opening/creating file: %v", err)
	}
	file.WriteString("Peer:\n")

	//Chiusura file
	defer file.Close()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
