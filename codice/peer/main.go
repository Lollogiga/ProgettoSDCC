package main

import (
	"codice/peer/api"
	"codice/peer/config"
	_ "codice/peer/config"
	"codice/peer/recovery"
	_ "codice/peer/registry"
	pb "codice/peer/registry"
	"codice/peer/service"
	"codice/peer/shared"
	"context"
	_ "encoding/json"
	"flag"
	"fmt"
	_ "fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// Flag per il controllo dei casi localhost o docker
	localhostFlag := flag.Bool("localhost", false, "Indica se il programma sta eseguendo su localhost")
	dockerFlag := flag.Bool("docker", false, "Indica se il programma sta eseguendo su Docker")
	flag.Parse()

	//Inizializzo il peer per eseguire in localhost / su docker
	if *localhostFlag {
		if flag.NArg() > 0 {
			config.MyAddress = flag.Arg(0)
		} else {
			fmt.Println("Usage: go run main.go -localhost IP:port")
			return
		}

		//Verifico che l'indirizzo sia corretto:
		api.VerifyAddress(config.MyAddress)

		//Prendo le informazioni dal file config.json
		config.LocalConfig()

	} else if *dockerFlag {
		log.Println("Docker environment")
		//Prendo le informazioni dal file config.json
		config.DockerConfiguration()

		//Verifico l'esistenza del file: "Id.yaml" --> Se esiste sono in crash-recovery:
		filePath := "recovery/Id.yaml"
		if recovery.FileExists(filePath) {
			log.Printf("I'm crash-recovery")
			recovery.RecoveryString = "Recovery"
			recovery.RecoveryId = recovery.GetId(filePath)
		}

	} else {
		fmt.Println("Specificare un flag: (-localHost) o (-docker)")
		return
	}

	//In tutti i casi mi connetto al peer registry
	log.Printf("The recoveryString is: %s, and the recoveryId %d", recovery.RecoveryString, recovery.RecoveryId)
	conn, err := grpc.Dial(config.ServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to registry: %v", err)
	}
	defer conn.Close()

	client := pb.NewRegistryClient(conn)

	// Invia una richiesta di join al registry
	resp, err := client.JoinNetwork(context.Background(), &pb.JoinRequest{Addr: config.MyAddress, RecoveryString: recovery.RecoveryString, RecoveryId: int32(recovery.RecoveryId)})
	if err != nil {
		log.Fatalf("failed to join network: %v", err)
	}

	//Estraggo risposta:
	log.Printf("Received ID: %d\n", resp.GetId())
	shared.MyId = resp.GetId()
	log.Println("Received Peer List:")

	for _, server := range resp.GetPeerList() {
		shared.NumNode++
		shared.PeerList = append(shared.PeerList, pb.PeerInfo{
			Id:     int32(server.GetId()),
			Addr:   server.GetAddr(),
			Leader: false,
		})
		log.Printf("ID: %d, Addr: %s", server.GetId(), server.GetAddr())
	}

	//In tutti i casi creo un nuovo file in cui inserisco il mio id(Se sono su docker)
	if *dockerFlag {
		recovery.SaveId(shared.MyId)
	}

	go startServer()

	if config.BullySelected == true {
		service.Bully()
	} else if config.DolevSelected == true {
		service.DolevStartElection()
	} else {
		shared.StimateLeader = shared.MyId
		service.DKR()
	}

	//Implementazione funzionalità peer:
	go api.Heartbeat(conn, err)

	for {

	}
}

func startServer() {
	// Crea un listener TCP
	lis, err := net.Listen("tcp", config.MyAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Crea un nuovo peer gRPC
	grpcServer := grpc.NewServer()

	// Registra i servizi sul peer gRPC
	registerServices(grpcServer)

	// Avvia il peer gRPC per gestire le richieste
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func registerServices(server *grpc.Server) {
	// Registra il servizio Time
	HeartBeatService := &service.HeartBeat{}
	pb.RegisterServiceServer(server, HeartBeatService)

	UpdateService := &service.UpdateServer{}
	pb.RegisterUpdateServer(server, UpdateService)

	Electionserver := &service.Election{}
	pb.RegisterElectionServer(server, Electionserver)

	//TODO registrare i restanti servizi

}
