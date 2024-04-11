package main

import (
	"codice/server/api"
	_ "codice/server/registry"
	pb "codice/server/registry" // Importa il pacchetto gRPC generato
	"codice/server/service"
	"codice/server/shared"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

/*

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}*/

func main() {

	// Ottieni gli argomenti dalla riga di comando
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalf("Usage: go run main.go IP:port")
	}
	shared.Address = args[0]

	//Verifico che l'indirizzo sia corretto:
	api.VerifyAddress(shared.Address)

	//Mi connetto al serverRegistry //TODO leggere da un file di configurazione
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to registry: %v", err)
	}
	defer conn.Close()

	client := pb.NewRegistryClient(conn)

	// Invia una richiesta di join al registry
	resp, err := client.JoinNetwork(context.Background(), &pb.JoinRequest{Addr: shared.Address})
	if err != nil {
		log.Fatalf("failed to join network: %v", err)
	}

	//Estraggo risposta:
	log.Printf("Received ID: %d\n", resp.GetId())
	shared.MyId = resp.GetId()
	log.Println("Received Server List:")

	for _, server := range resp.GetPeerList() {
		shared.NumNode++
		shared.PeerList = append(shared.PeerList, pb.PeerInfo{
			Id:     int32(server.GetId()),
			Addr:   server.GetAddr(),
			Leader: false,
		})
		log.Printf("ID: %d, Addr: %s", server.GetId(), server.GetAddr())
	}

	go startServer()

	//service.Bully()
	service.DolevStartElection()

	//Avvio procedura per tenere traccia degli aggiornamenti
	//go RegisterService()
	//Implementazione funzionalità peer:
	go api.GetTime(conn, err)

	for {
	}
}

func startServer() {
	// Crea un listener TCP
	lis, err := net.Listen("tcp", shared.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Crea un nuovo server gRPC
	grpcServer := grpc.NewServer()

	// Registra i servizi sul server gRPC
	registerServices(grpcServer)

	// Avvia il server gRPC per gestire le richieste
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func registerServices(server *grpc.Server) {
	// Registra il servizio Time
	timeService := &service.Time{}
	pb.RegisterServiceServer(server, timeService)

	UpdateService := &service.UpdateServer{}
	pb.RegisterUpdateServer(server, UpdateService)

	Electionserver := &service.Election{}
	pb.RegisterElectionServer(server, Electionserver)

	//TODO registrare i restanti servizi

}

/*func main() {
filePath := "/percorso/del/tuo/file"
if fileExists(filePath) {
	fmt.Printf("Il file %s esiste.\n", filePath)
} else {
	fmt.Printf("Il file %s non esiste.\n", filePath)
}
*/
