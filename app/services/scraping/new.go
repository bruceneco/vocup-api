package services

const (
	baseURL = "https://www.dicio.com.br"
)

type ScrapperService interface {
	GetMostSearchedWords(int) []string
}

type scrapperService struct {
}

func NewScrapperService() ScrapperService {
	return &scrapperService{}
}
