package scrapper

import (
	"github.com/gocolly/colly"
)

func Run() string {

	c := colly.NewCollector(
		colly.AllowedDomains("dolarhoy.com"),
	)

	valordolar := ""

	c.OnHTML(".is-5", func(e *colly.HTMLElement) {
		valordolar = e.ChildText(".val")
	})

	c.Visit("https://dolarhoy.com/")

	return valordolar
}