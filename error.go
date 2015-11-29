package scraper

import "fmt"

type ScraperError struct {
	Proxy       string
	URL         string
	ErrorString string
}

func (e ScraperError) Error() string {
	return fmt.Sprintf("Error get %s: %s", e.URL, e.ErrorString)
}
