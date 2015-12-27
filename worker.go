package scraper

import (
	"net/http"
	"net/url"
)

func (s *Scraper) runWorker() {
	defer func() {
		s.wg.Done()
		<-s.concurrentCounter
	}()
	var httpClient *http.Client
	var currentProxy *url.URL
	for task := range s.queue {
	page:
		for try := 0; try < s.maxRetry; try++ {
			if len(s.proxyList) > 0 {
				currentProxy = s.getNextProxy()
				httpClient = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(currentProxy)}}
			} else {
				currentProxy = nil
				httpClient = &http.Client{}
			}
			resp, err := httpClient.Get(task)
			if err != nil {
				errScraper := ScraperError{currentProxy, task, err.Error()}
				result := Response{http.Response{}, errScraper, currentProxy}
				s.Results <- result
			} else if resp.StatusCode >= 300 {
				errScraper := ScraperError{currentProxy, task, resp.Status}
				result := Response{*resp, errScraper, currentProxy}
				s.Results <- result
			} else {
				result := Response{*resp, nil, currentProxy}
				s.Results <- result
				break page
			}
		}
	}
}
