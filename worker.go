package scraper

import (
	"net/http"
	"net/url"
	"time"
)

func (s *Scraper) runWorker() {
	defer func() {
		s.wg.Done()
		<-s.concurrentCounter
	}()
	var httpClient *http.Client
	var currentProxy *url.URL
	for s.notDone {
		select {
		case task := <-s.queue:
		page:
			for try := 0; try < s.maxRetry; try++ {
				if len(s.proxyList) > 0 {
					currentProxy = s.getNextProxy()
					httpClient = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(currentProxy)}}
				} else {
					httpClient = &http.Client{}
				}
				resp, err := httpClient.Get(task)
				if err != nil {
					errScraper := ScraperError{currentProxy.String(), task, err.Error()}
					s.patern.LogError(errScraper)
				} else if resp.StatusCode > 300 {
					errScraper := ScraperError{currentProxy.String(), task, resp.Status}
					s.patern.LogError(errScraper)
				} else {
					s.patern.Parse(resp)
					s.patern.Save()
					break page
				}
			}
		case <-time.After(time.Second * 1):
			return
		}
	}
}
