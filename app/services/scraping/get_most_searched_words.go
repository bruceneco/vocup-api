package services

import (
	"github.com/gocolly/colly"
	"log"
	"strings"
	"sync"
)

func (*scrapperService) GetMostSearchedWords(nPages int) []string {
	var mostSearchedWords []string
	mu := sync.Mutex{}

	var wg sync.WaitGroup
	wg.Add(nPages)
	for i := 0; i < nPages; i++ {
		i := i
		go func() {
			c := colly.NewCollector()

			c.OnHTML(".list > li", func(e *colly.HTMLElement) {
				mu.Lock()
				mostSearchedWords = append(mostSearchedWords, strings.TrimSpace(e.DOM.Text()))
				mu.Unlock()
			})
			c.OnScraped(func(response *colly.Response) {
				wg.Done()
			})
			c.OnError(func(r *colly.Response, err error) {
				log.Printf("%v", err)
				wg.Done()
			})
			err := c.Visit(strings.Join([]string{baseURL, "palavras-mais-buscadas"}, "/"))
			if err != nil {
				log.Printf("error scrapping page %d\n", i)
			}
		}()
	}
	wg.Wait()
	return mostSearchedWords
}
