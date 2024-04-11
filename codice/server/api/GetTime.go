package api

import (
	"codice/server/registry"
	"codice/server/service"
	shared "codice/server/shared"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

// Funzionalità peer che si occupa di Richiedere l'ora esatta
func GetTime(conn *grpc.ClientConn, err error) {
	for {
		log.Printf("Request Time:\n")
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

				timeResp, leaderror := peer.GetTime(context.Background(), &registry.TimeRequest{
					Message: "TIME",
				})
				if leaderror != nil {
					log.Printf("Leader unreachable: \n")
					//service.Bully()
					service.DolevStartElection()
				} else {
					log.Printf("Leader response: \nTime: %s", timeResp.Message)
				}
			}
		}
	}
}
