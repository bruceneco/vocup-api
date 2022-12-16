package main

import (
	"github.com/bruceneco/vocup-api/api/scrap"
	"github.com/bruceneco/vocup-api/app/server/scrap"
	"github.com/bruceneco/vocup-api/app/services"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main() {

	// Services
	log := hclog.Default()
	scrapperService := services.NewScrapperService()
	scrapServer := server.NewScrapServer(scrapperService)

	// Servers
	gs := grpc.NewServer()
	scrap.RegisterScrapServer(gs, scrapServer)

	// Application start
	reflection.Register(gs)
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Error("error on listening to port 8080", "error", err)
		os.Exit(1)
	}
	log.Info("server running", "port", lis.Addr())
	if err = gs.Serve(lis); err != nil {
		log.Error("fail on serve", "error", err)
	}
}
