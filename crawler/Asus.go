package crawler

import (
	"fmt"
	"strings"

	"github.com/gen2brain/beeep"
	colly "github.com/gocolly/colly/v2"
)

func Asus() {
	fmt.Println("--- Initiating Asus crawler ---")

	c := colly.NewCollector()

	var name string = ""
	var price string = ""

	c.OnHTML("h1.product-detail-name", func(e *colly.HTMLElement) {
		name = strings.TrimSpace(e.Text)
	})

	c.OnHTML("p.product-detail-price", func(e *colly.HTMLElement) {
		price = strings.TrimSpace(e.Text)
	})

	c.OnHTML("button.btn-buy", func(e *colly.HTMLElement) {
		fmt.Printf("[Asus] Available: %s for %s\n", name, price)
		var message = fmt.Sprintf("[Asus] Available for %s", price)
		err := beeep.Alert(name, message, "icon.jpg")
		if err != nil {
			panic(err)
		}
	})

	c.OnHTML("button.btn-notify", func(e *colly.HTMLElement) {
		fmt.Printf("[Asus] Not available: %s for %s\n", name, price)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(err)
	})

	c.Visit("https://webshop.asus.com/ch-en/90YV0IE1-M0NA00/TUF-RTX4090-24G-GAMING")
	c.Visit("https://webshop.asus.com/ch-en/90YV0IE0-M0NA00/TUF-RTX4090-O24G-GAMING")
	c.Visit("https://webshop.asus.com/ch-en/90YV0ID0-M0NA00/ASUS-ROG-Strix-GeForce-RTX-4090-24GB-OC-Edition-Gaming-Grafikkarte")

	fmt.Println("--- Asus crawling finished ---")
}
