package main

import (
	"github.com/bruceneco/vocup-api/api/scrap"
	"github.com/bruceneco/vocup-api/app/repository"
	"github.com/bruceneco/vocup-api/app/server/scrap"
	"github.com/bruceneco/vocup-api/app/services/scraping"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func main() {
	// Database and DAO
	dao := repository.NewDAO()
	_, err := repository.NewDB()
	if err != nil {
		log.Fatalf("cant create database: %v", err)
	}
	// Services
	log := hclog.Default()
	scrapperService := services.NewScrapperService(dao)
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
