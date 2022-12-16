package main

import (
	"fmt"
	"github.com/bruceneco/vocup-api/app/services"
	usecases "github.com/bruceneco/vocup-api/app/usecases/scrap"
)

func main() {
	scrapperService := services.NewScrapperService()
	scrapUC := usecases.NewScrapUsecase(scrapperService)

	fmt.Println(scrapUC.ScrapTopWords(), len(scrapUC.ScrapTopWords()))
}
