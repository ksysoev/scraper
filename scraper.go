package scraper

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"sync"
)

//Scraper - structure for crawler, dont use manual create object, use NewScraper(int(maximum concarensy), Patern(output format and functions for parsing))
type Scraper struct {
	proxyList         []*url.URL
	linksList         []string
	maxConcurrent     int
	patern            Patern
	getNextProxy      func() *url.URL
	wg                sync.WaitGroup
	concurrentCounter chan bool
}

//Patern interface must be use to give crawler format output data and function for parsing content and function to save data. Data type  must have methods Parse(io.Reader) and Save().
type Patern interface {
	Parse(io.Reader)
	Save()
}

//NewScraper - use for create new  crawler.
func NewScraper(maxConcurrent int, patern Patern) *Scraper {
	s := Scraper{maxConcurrent: maxConcurrent}
	s.getNextProxy = s.nextProxy()
	s.patern = patern
	s.concurrentCounter = make(chan bool, s.maxConcurrent)
	return &s
}

//RunCrawler - use this method for start crawler, after add links and proxy(if need).
func (s *Scraper) RunCrawler() {
	s.wg = sync.WaitGroup{}

	for currenLinkID := range s.linksList {
		s.wg.Add(1)
		go s.getPage(s.linksList[currenLinkID])
		s.concurrentCounter <- true

	}
	s.wg.Wait()
}

func (s *Scraper) getPage(href string) {
	defer s.wg.Done()
	var httpClient *http.Client
	if len(s.proxyList) > 0 {
		httpClient = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(s.getNextProxy())}}
	} else {
		httpClient = &http.Client{}
	}
	resp, err := httpClient.Get(href)
	if err != nil {
		fmt.Println(err)
	} else {
		s.patern.Parse(resp.Body)
		s.patern.Save()
	}
	// patern := s.patern
	<-s.concurrentCounter
}

//AddProxy - function for adding address's of proxy servers to crawler, address must be in format:
// http(s)://ip.or.domain.name:port
func (s *Scraper) AddProxy(proxy string) error {
	urlValidator := regexp.MustCompile(`^https?:\/\/([\da-z\.-]+)\.[\da-z\.]{2,6}:\d+$`)
	if !urlValidator.MatchString(proxy) {
		return errors.New("Proxy does not valid url, format for proxy  sheme://address.or.dns:port")
	}
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return err
	}
	s.proxyList = append(s.proxyList, proxyURL)
	return nil
}

func (s *Scraper) nextProxy() func() *url.URL {
	curentProxyID := 0
	return func() *url.URL {
		if len(s.proxyList) > 0 {
			if curentProxyID >= len(s.proxyList) {
				curentProxyID = 0
			}
			proxy := s.proxyList[curentProxyID]
			curentProxyID++
			return proxy
		}
		return nil
	}
}

//AddLink -  function to adding URL for parcing. Function have one argument url in string format.
func (s *Scraper) AddLink(link string) error {
	_, err := url.Parse(link)
	if err != nil {
		return err
	}
	s.linksList = append(s.linksList, link)
	return nil

}
