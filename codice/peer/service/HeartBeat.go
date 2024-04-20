package service

import (
	pb "codice/peer/registry"
	"context"
	"log"
)

type HeartBeat struct {
	pb.UnimplementedServiceServer
}

func (s *HeartBeat) HeartBeat(ctx context.Context, req *pb.HeartBeatMessage) (*pb.HeartBeatMessage, error) {
	resultChan := make(chan *pb.HeartBeatMessage)
	log.Printf("I'm alive\n")
	go func() {
		if req.Message == "HeartBeat" {
			reply := &pb.HeartBeatMessage{Message: "Alive"}
			resultChan <- reply
		}
	}()
	reply := <-resultChan
	return reply, nil
}
