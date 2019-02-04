package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Lighty0410/ekadashi-server/pkg/mongo"
	"github.com/Lighty0410/ekadashi-server/pkg/server"
)

func main() {
	mongoService, err := mongo.NewService()
	if err != nil {
		log.Fatalf("Could not create mongo service: %v", err)
	}
	ekadashiServer, err := server.NewEkadashiServer(mongoService)
	if err != nil {
		log.Fatalf("Could not create ekadashi server: %v", err)
	}
	server := &http.Server{
		Addr:    ":9000",
		Handler: ekadashiServer,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("Could not listen: %v", err)
		}
	}()

	sig := <-stop

	log.Printf("Shutting down due to signal: %v", sig)
	err = server.Shutdown(context.Background())
	if err != nil {
		log.Fatalf("Cannot shutdown the server")
	}
}
