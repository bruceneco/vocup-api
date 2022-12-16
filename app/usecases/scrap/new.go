package usecases

import "github.com/bruceneco/vocup-api/app/services"

type ScrapUsecase interface {
	ScrapTopWords() []string
}

type scrapUsecase struct {
	scrapperService services.ScrapperService
}

func NewScrapUsecase(scrapperService services.ScrapperService) ScrapUsecase {
	return &scrapUsecase{scrapperService: scrapperService}
}

func (s *scrapUsecase) ScrapTopWords() []string {
	return s.scrapperService.GetMostSearchedWords(50)
}
