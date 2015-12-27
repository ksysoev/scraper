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
	getNextProxy      func() *url.URL
	wg                sync.WaitGroup
	concurrentCounter chan bool
	queue             chan string
	Results           chan Response
	notDone           bool
}

type Response struct {
	http.Response
	Err   error
	Proxy *url.URL
}

//NewScraper - use for create new  crawler.
func NewScraper(maxConcurrent int, maxRetry int) *Scraper {
	s := Scraper{maxConcurrent: maxConcurrent, maxRetry: maxRetry}
	s.getNextProxy = s.nextProxy()
	s.concurrentCounter = make(chan bool, s.maxConcurrent)
	s.queue = make(chan string)
	s.Results = make(chan Response)
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
	close(s.Results)
}
