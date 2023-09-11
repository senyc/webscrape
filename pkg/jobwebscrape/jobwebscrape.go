package jobwebscrape

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func SearchJobBoard() {
	url := "https://www.monster.com/jobs/search?q=javascript+developer"
	c := colly.NewCollector()

	c.OnHTML("ul", func(e *colly.HTMLElement) {
		e.ForEach("li", func(index int, el *colly.HTMLElement) {
			el.ForEach("div", func(index int, d *colly.HTMLElement) {
				d.ForEach("h3", func(index int, t *colly.HTMLElement) {
					fmt.Print(t.Text)
				})
			})
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
