package server

import (
	pb "github.com/bruceneco/vocup-api/api/scrap"
	"github.com/bruceneco/vocup-api/app/services"
)

type scrapServer struct {
	pb.ScrapServer
	scrapperService services.ScrapperService
}

func NewScrapServer(scrapperService services.ScrapperService) pb.ScrapServer {
	return &scrapServer{scrapperService: scrapperService}
}
