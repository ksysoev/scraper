# scraper
Simple go web scraper

#Example

  type PaternTest struct {
	   Urls map[string]string
   }

   func (p *PaternRubrics) Parse(resp *http.Response) {
	    doc, err := goquery.NewDocumentFromReader(resp.Body)
	     if err != nil {
		       fmt.Println(err)
	          } else {
		            p.Urls = make(map[string]string)
		              doc.Find("a").Each(func(i int, item *goquery.Selection) {
			                 name := item.Text()
			                    href, _ := item.Attr("href")
			                       p.Urls[href] = name
		                           })
	                            }
                            }

                            func (p *PaternRubrics) Save() {
	                             for key, value := range p.Urls {
                                 fmt.Printf("%s - %s", key, value)
	                              }
                              }

                              scrap := scraper.NewScraper(1, 3, &TestPatern{})
                                scrap.AddProxy("http://testproxy.ru:8080")
                                scrap.AddLink("http://testpage.ru/")
                                scrap.RunCrawler()
