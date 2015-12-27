package scraper

import "testing"

func TestNewScraper(t *testing.T) {
	s := NewScraper(1, 1)
	if s.maxConcurrent != 1 {
		t.Error("Can't create a Scraper")
	}
}

func TestRunScraper(t *testing.T) {
	s := NewScraper(1, 1)
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
