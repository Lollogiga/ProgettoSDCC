package main

import (
	pb "codice/server_registry/registry"
	"codice/server_registry/service"
	"codice/server_registry/shared"
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	// Flag per il controllo dei casi localhost o docker
	shared.Localhostflag = flag.Bool("localhost", false, "Indica se il programma sta eseguendo su localhost")
	shared.DockerFlag = flag.Bool("docker", false, "Indica se il programma sta eseguendo su Docker")
	flag.Parse()
	if !*shared.DockerFlag && !*shared.Localhostflag {
		log.Fatalf("Use flag -localhost or -docker")
	}

	server := &service.RegistryServer{
		PeerList: []*pb.PeerInfo{},
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRegistryServer(s, server)
	log.Println("Registry listening on port 50051")

	shared.File, err = os.OpenFile("peerList/PeerList.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	//file, err = os.OpenFile("peerList.yaml", os.O_RDWR, 0600)

	if err != nil {
		log.Fatalf("error opening/creating file: %v", err)
	}
	_, err = shared.File.WriteString("Peer:\n")
	if err != nil {
		return
	}

	//Chiusura file
	defer shared.File.Close()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
