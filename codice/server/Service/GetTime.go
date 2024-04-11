package service

import (
	pb "codice/server/registry"
	"context"
	"log"
	"time"
)

type Time struct {
	pb.UnimplementedServiceServer
}

func (s *Time) GetTime(ctx context.Context, req *pb.TimeRequest) (*pb.TimeReply, error) {
	// Implementazione della funzione GetTime del servizio Time
	// Crea un canale per ricevere il risultato dalla goroutine
	resultChan := make(chan *pb.TimeReply)
	log.Printf("I'm alive\n")
	// Esegui la logica all'interno di una goroutine
	go func() {
		// Ottieni l'ora corrente
		currentTime := time.Now()

		// Costruisci la risposta
		reply := &pb.TimeReply{
			Message: currentTime.Format(time.RFC3339), // Formatta l'ora corrente come stringa
		}

		// Invia la risposta attraverso il canale
		resultChan <- reply
	}()

	// Ricevi il risultato dalla goroutine
	reply := <-resultChan

	return reply, nil
}
