package scraper

import (
	"fmt"
	"net/url"
)

type ScraperError struct {
	Proxy       *url.URL
	URL         string
	ErrorString string
}

func (e ScraperError) Error() string {
	return fmt.Sprintf("Error get %s: %s", e.URL, e.ErrorString)
}
