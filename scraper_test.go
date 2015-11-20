package scraper

import (
	"net/http"
	"testing"
)

var ParseOk = true
var SaveOk = true

type PaternStructure struct {
	test string
}

func (p *PaternStructure) Parse(resp *http.Response) {
	ParseOk = false
}

func (p *PaternStructure) Save() {
	SaveOk = false
}

func TestNewScraper(t *testing.T) {
	testPatern := PaternStructure{}
	s := NewScraper(1, 1, &testPatern)
	if s.maxConcurrent != 1 {
		t.Error("Can't create a Scraper")
	}
}

func TestRunScraper(t *testing.T) {
	testPatern := PaternStructure{}
	s := NewScraper(1, 1, &testPatern)
	s.RunCrawler()
}

// func TestGetPage(t *testing.T) {
//
// 	ParseOk = true
// 	SaveOk = true
// 	testPatern := PaternStructure{}
// 	s := NewScraper(1, &testPatern)
// 	s.getPage("http://127.0.0.1")
// 	if ParseOk || SaveOk {
// 		t.Error("Error get page")
// 	}
// }
