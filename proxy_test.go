package scraper

import (
	"net/url"
	"testing"
)

func TestAddProxy(t *testing.T) {
	s := Scraper{}
	s.AddProxy("http://127.0.0.1:8080")
	proxyURL, _ := url.Parse("http://127.0.0.1:8080")
	if s.proxyList[0].url.String() != proxyURL.String() {
		t.Error("Error add proxy server")
	}
}

func TestAddProxyFail(t *testing.T) {
	s := Scraper{}
	err := s.AddProxy("http!//Not_a_Proxy_Address:70000")
	if err == nil {
		t.Error("Error check proxy url")
	}
}

func TestNextProxy(t *testing.T) {
	s := Scraper{}

	getNextProxyUrl := s.nextProxy()

	proxyURL1, _ := url.Parse("http://127.0.0.1:8080")
	proxyURL2, _ := url.Parse("http://127.0.0.2:8080")

	if getNextProxyUrl() != nil {
		t.Error("Error get next proxy")
	}

	s.AddProxy("http://127.0.0.1:8080")
	s.AddProxy("http://127.0.0.2:8080")

	if getNextProxyUrl().url.String() != proxyURL1.String() {
		t.Error("Error get next proxy")
	}

	if getNextProxyUrl().url.String() != proxyURL2.String() {
		t.Error("Error get next proxy")
	}

	if getNextProxyUrl().url.String() != proxyURL1.String() {
		t.Error("Error get next proxy")
	}
}

func TestRemoveFailedProxy(t *testing.T) {
	s := NewScraper(1, 1)
	s.AddProxy("http://127.0.0.1:8080")

	proxy := s.getNextProxy()
	for i := 0; i < 10; i++ {
		proxy.gotFail()
	}

	if s.getNextProxy() != nil {
		t.Error("Error remove a failed proxy")
	}

}

func TestRemoveFailedProxyMulti(t *testing.T) {
	s := NewScraper(1, 1)
	s.AddProxy("http://127.0.0.1:8080")
	s.AddProxy("http://127.0.0.2:8080")

	proxyURL2, _ := url.Parse("http://127.0.0.2:8080")

	proxy := s.getNextProxy()
	for i := 0; i < 10; i++ {
		proxy.gotFail()
	}

	if s.getNextProxy().url.String() != proxyURL2.String() {
		t.Error("Got wrong proxy address")
	}

	if s.getNextProxy().url.String() != proxyURL2.String() {
		t.Error("Got wrong proxy address")
	}

	proxy = s.getNextProxy()
	for i := 0; i < 10; i++ {
		proxy.gotFail()
	}

	if s.getNextProxy() != nil {
		t.Error("Error remove a failed proxy")
	}
}
