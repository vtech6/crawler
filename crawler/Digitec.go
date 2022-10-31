package crawler

import (
	"fmt"

	"github.com/gen2brain/beeep"
	colly "github.com/gocolly/colly/v2"
)

func Digitec() {
	c := colly.NewCollector()
	fmt.Println("!!!!!Initiating Digitec Crawler!!!!!")
	var disabled bool
	var price string = ""
	var name string
	// Find and visit all links
	c.OnHTML("h1.sc-12r9jwk-0.kFGyqi", func(e *colly.HTMLElement) {
		name = e.Text
	})
	c.OnHTML("strong.sc-1ttlt4t-1.kRYPrq", func(e *colly.HTMLElement) {
		price = e.Text
	})

	c.OnHTML("button.sc-1olg58b-0.cXWAZg.sc-185g6wq-6.bvMIBV", func(e *colly.HTMLElement) {
		_, disabled = e.DOM.Attr("disabled")
		if disabled || price == "" {
			fmt.Printf("Not available: %s\n", name)
			price = ""
		} else {
			fmt.Printf("Available: %s for %s\n", name, price)
			var message = fmt.Sprintf("Available for %s", price)
			err := beeep.Alert(name, message, "icon.jpg")
			if err != nil {
				panic(err)
			}
			price = ""
		}
	})

	c.OnRequest(func(r *colly.Request) {
		// time.Sleep(7 * time.Second)
	})

	c.Visit("https://www.digitec.ch/en/s1/product/asus-tuf-gaming-geforce-rtx-4090-oc-edition-24-gb-graphics-card-22556792")
	c.Visit("https://www.digitec.ch/en/product/gigabyte-giga-vga-24gb-rtx4090-aorus-master-24g-3xdp3xhdmi-aorus-geforce-rtx-4090-master-24g-24-gb-g-22822881")
	c.Visit("https://www.digitec.ch/en/product/msi-geforce-rtx-4090-suprim-x-24g-24-gb-graphics-card-22532748")
	c.Visit("https://www.digitec.ch/en/product/gigabyte-geforce-rtx-4090-gaming-oc-24g-24-gb-graphics-card-22594251")
	c.Visit("https://www.digitec.ch/en/product/msi-geforce-rtx-4090-gaming-x-trio-24g-24-gb-graphics-card-22534974")
	c.Visit("https://www.digitec.ch/en/product/asus-rog-strix-geforce-rtx-4090-oc-edition-24-gb-graphics-card-22557847")
	c.Visit("https://www.digitec.ch/en/product/asus-rog-strix-geforce-rtx-4090-24-gb-graphics-card-22662082")
	c.Visit("https://www.digitec.ch/en/product/msi-geforce-rtx-4090-gaming-x-trio-24g-24-gb-graphics-card-22534974")
	c.Visit("https://www.digitec.ch/en/product/asus-tuf-gaming-geforce-rtx-4090-24-gb-graphics-card-22662761")
	c.Visit("https://www.digitec.ch/en/product/msi-geforce-rtx-4090-gaming-trio-24g-24-gb-graphics-card-22666076")
	c.Visit("https://www.digitec.ch/en/product/gigabyte-giga-vga-24gb-rtx4090-windforce-24g-3xdp3xhdmi-geforce-rtx-4090-windforce-24g-24-gb-graphic-22822880")

	fmt.Println("!!!!!Digitec crawling finished!!!!!\n")
}
