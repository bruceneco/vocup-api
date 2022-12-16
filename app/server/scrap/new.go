package server

import (
	"context"
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

func (s *scrapServer) GetTopWords(_ context.Context, req *pb.GetTopWordsReq) (*pb.GetTopWordsRes, error) {
	return &pb.GetTopWordsRes{
		Words: s.scrapperService.GetMostSearchedWords(int(req.NumPages)),
	}, nil
}
