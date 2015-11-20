package scraper

import (
	"net/url"
	"testing"
)

func TestAddProxy(t *testing.T) {
	s := Scraper{}
	s.AddProxy("http://127.0.0.1:8080")
	proxyUrl, _ := url.Parse("http://127.0.0.1:8080")
	if s.proxyList[0].String() != proxyUrl.String() {
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

	proxyUrl1, _ := url.Parse("http://127.0.0.1:8080")
	proxyUrl2, _ := url.Parse("http://127.0.0.2:8080")

	if getNextProxyUrl() != nil {
		t.Error("Error get next proxy")
	}

	s.AddProxy("http://127.0.0.1:8080")
	s.AddProxy("http://127.0.0.2:8080")

	if getNextProxyUrl().String() != proxyUrl1.String() {
		t.Error("Error get next proxy")
	}

	if getNextProxyUrl().String() != proxyUrl2.String() {
		t.Error("Error get next proxy")
	}

	if getNextProxyUrl().String() != proxyUrl1.String() {
		t.Error("Error get next proxy")
	}
}
