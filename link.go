package scraper

import "net/url"

//AddLink -  function to adding URL for parcing. Function have one argument url in string format.
func (s *Scraper) AddLink(link string) error {
	_, err := url.Parse(link)
	if err != nil {
		return err
	}
	s.linksList = append(s.linksList, link)
	return nil

}
