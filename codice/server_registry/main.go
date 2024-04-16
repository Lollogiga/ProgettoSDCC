package main

import (
	pb "codice/server_registry/registry"
	"codice/server_registry/service"
	"codice/server_registry/shared"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

//var PeerList []pb.PeerInfo

func main() {

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

	shared.File, err = os.OpenFile("peerList/peerList.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
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
