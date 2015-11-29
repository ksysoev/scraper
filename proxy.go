package scraper

import (
	"errors"
	"net/url"
	"regexp"
)

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
