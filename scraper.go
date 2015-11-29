package scraper

import (
	"net/http"
	"net/url"
	"sync"
)

//Scraper - structure for crawler, dont use manual create object, use NewScraper(int(maximum concarensy), Patern(output format and functions for parsing))
type Scraper struct {
	proxyList         []*url.URL
	linksList         []string
	maxConcurrent     int
	maxRetry          int
	patern            Patern
	getNextProxy      func() *url.URL
	wg                sync.WaitGroup
	concurrentCounter chan bool
	queue             chan string
	notDone           bool
}

//Patern interface must be use to give crawler format output data and function for parsing content and function to save data. Data type  must have methods Parse(io.Reader) and Save().
type Patern interface {
	Parse(*http.Response)
	Save()
	LogError(error)
}

//NewScraper - use for create new  crawler.
func NewScraper(maxConcurrent int, maxRetry int, patern Patern) *Scraper {
	s := Scraper{maxConcurrent: maxConcurrent, maxRetry: maxRetry}
	s.getNextProxy = s.nextProxy()
	s.patern = patern
	s.concurrentCounter = make(chan bool, s.maxConcurrent)
	s.queue = make(chan string)
	s.notDone = true
	return &s
}

//RunCrawler - use this method for start crawler, after add links and proxy(if need).
func (s *Scraper) RunCrawler() {
	s.wg = sync.WaitGroup{}
	go s.runQueue()
	for s.notDone {
		go s.runWorker()
		s.wg.Add(1)
		s.concurrentCounter <- true
	}
	s.wg.Wait()
}
