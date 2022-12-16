package services

import "github.com/bruceneco/vocup-api/app/repository"

const (
	baseURL = "https://www.dicio.com.br"
)

type ScrapperService interface {
	GetMostSearchedWords(int) []string
}

type scrapperService struct {
	dao repository.DAO
}

func NewScrapperService(dao repository.DAO) ScrapperService {
	return &scrapperService{dao: dao}
}
