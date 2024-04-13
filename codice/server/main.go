package main

import (
	"codice/server/api"
	"codice/server/config"
	_ "codice/server/config"
	_ "codice/server/registry"
	pb "codice/server/registry" // Importa il pacchetto gRPC generato
	"codice/server/service"
	"codice/server/shared"
	"context"
	_ "encoding/json"
	"flag"
	"fmt"
	_ "fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

/*

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}*/

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
		//Prendo le informazioni dal file config.json
		config.DockerConfiguration()
	} else {
		fmt.Println("Specificare un flag: (-localHost) o (-docker)")
		return
	}

	//TODO gestire nel caso docker il ravvio automatico (creare un file in cui inserisco il mio ID)

	//Mi connetto al serverRegistry:
	conn, err := grpc.Dial(config.ServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to registry: %v", err)
	}
	defer conn.Close()

	client := pb.NewRegistryClient(conn)

	// Invia una richiesta di join al registry
	resp, err := client.JoinNetwork(context.Background(), &pb.JoinRequest{Addr: config.MyAddress})
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

	if config.BullySelected == true {
		service.Bully()
	} else {
		service.DolevStartElection()
	}

	//Implementazione funzionalità peer:
	go api.GetTime(conn, err)

	for {
	}
}

func startServer() {
	// Crea un listener TCP
	lis, err := net.Listen("tcp", config.MyAddress)
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
