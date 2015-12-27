package scraper

func (s *Scraper) runQueue() {
	for _, currenLink := range s.linksList {
		s.queue <- currenLink
	}
	s.notDone = false
	close(s.queue)
}
