package server

import (
	"context"
	"github.com/bruceneco/vocup-api/api/scrap"
)

func (s *scrapServer) GetTopWords(_ context.Context, req *scrap.GetTopWordsReq) (*scrap.GetTopWordsRes, error) {
	return &scrap.GetTopWordsRes{
		Words: s.scrapperService.GetMostSearchedWords(int(req.NumPages)),
	}, nil
}
