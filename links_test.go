package scraper

import "testing"

func TestAddLink(t *testing.T) {
	s := Scraper{}
	s.AddLink("http://127.0.0.1")
	if s.linksList[0] != "http://127.0.0.1" {
		t.Error("Error add link server")
	}
}
