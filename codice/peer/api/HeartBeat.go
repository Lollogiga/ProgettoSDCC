package api

import (
	"codice/peer/config"
	"codice/peer/registry"
	"codice/peer/service"
	shared "codice/peer/shared"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func Heartbeat(conn *grpc.ClientConn, err error) {
	for {
		log.Printf("Try to connect to leader(Id %d)\n", shared.LeaderId)
		time.Sleep(5000 * time.Millisecond)
		for _, leader := range shared.PeerList {
			if leader.Leader == true {
				shared.LeaderId = leader.Id
				conn, err = grpc.Dial(leader.Addr, grpc.WithInsecure())
				if err != nil {
					log.Fatalf("failed to connect to registry: %v", err)
				}
				defer conn.Close()

				peer := registry.NewServiceClient(conn)

				_, leaderror := peer.HeartBeat(context.Background(), &registry.HeartBeatMessage{Message: "HeartBeat"})
				if leaderror != nil {
					log.Printf("Leader unreachable: \n")
					if config.BullySelected == true {
						service.Bully()
					} else if config.DKRSelected == true {
						service.DKR()
					} else {
						service.DolevStartElection()
					}

				} else {
					log.Printf("The leader is alive\n")
				}
			}
		}
	}
}
