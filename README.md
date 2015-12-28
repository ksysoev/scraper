# scraper
Simple go web scraper

#Example

```go
import "github.com/pestkam/scraper"

scrap := scraper.NewScraper(1, 3)
scrap.AddLink("http://example.com")
go scrap.RunCrawler()
for result := range scrap.Results {
  if result.Err != nil {
    fmt.Println(result.Err)
    continue
  }
  body, _ := ioutil.ReadAll(result.Body)
  fmt.Println(string(body))
}
```
