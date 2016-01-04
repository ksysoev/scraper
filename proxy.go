package scraper

import (
	"errors"
	"net/url"
	"regexp"
)

type proxyItem struct {
	url   *url.URL
	fails int
	uses  int
}

//AddProxy - function for adding address's of proxy servers to crawler, address must be in format:
// http(s)://ip.or.domain.name:port
func (s *Scraper) AddProxy(proxy string) error {
	urlValidator := regexp.MustCompile(`^https?:\/\/([\da-z\.-]+)\.[\da-z\.]{2,6}:\d+$`)
	if !urlValidator.MatchString(proxy) {
		return errors.New("Proxy does not valid url, format for proxy  sheme://address.or.domain:port")
	}
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return err
	}
	Proxy := proxyItem{url: proxyURL, fails: 0, uses: 0}
	s.proxyList = append(s.proxyList, Proxy)
	return nil
}

func (s *Scraper) nextProxy() func() *proxyItem {
	curentProxyID := 0
	return func() *proxyItem {
		if len(s.proxyList) > 0 {
			if curentProxyID >= len(s.proxyList) {
				curentProxyID = 0
			}
			proxy := &s.proxyList[curentProxyID]
			curentProxyID++
			if !proxy.isOk() {
				s.proxyList = append(s.proxyList[:curentProxyID-1], s.proxyList[curentProxyID:]...)
				if len(s.proxyList) < 1 {
					return nil
				}
				proxy = s.getNextProxy()
			}
			return proxy
		}
		return nil
	}
}

func (p *proxyItem) gotFail() {
	p.fails++
	p.uses++
}

func (p *proxyItem) gotSuccess() {
	p.uses++
}

func (p *proxyItem) isOk() bool {
	if p.fails <= 5 {
		return true
	}

	percentSuccess := (p.uses - p.fails) * 100 / p.uses
	if percentSuccess > 60 {
		return true
	}
	return false
}
