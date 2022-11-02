package crawler

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gen2brain/beeep"
	colly "github.com/gocolly/colly/v2"
)

func Smeg() {
	c := colly.NewCollector()
	fmt.Println("!!!!!Initiating Smeg Crawler!!!!!")
	var disabled bool
	var price string
	var name string

	wantedCards :=strings.Split(os.Getenv("WANTED_CARDS"), " ")

	// Find and visit all links
	c.OnHTML("#content", func(e *colly.HTMLElement) {
		e.ForEach("div.productGridElement.lg-flex-item.sm-basis-full.lg-basis-25.product-element", func(ind int, subE *colly.HTMLElement) {

			name = strings.TrimSpace(subE.ChildText("h2"))
			var err error
			price = subE.ChildText("div.generalPrice")
			var formattedPrice int
			formattedPrice, err = strconv.Atoi(strings.Replace(strings.Replace(strings.TrimSpace(price), ".", "", -1), "'", "", -1))
			if err != nil {
				panic(err)
			}

			disabled = subE.ChildAttr("i.fas.fa-truck.avaIcon.iconShipment", "style") == "color:#BFBFBF"
			if formattedPrice < 100000 || formattedPrice > 300000 {
				return
			}
			if !disabled {
				fmt.Printf("Available: %s for %s\n", name, price)
				var message = fmt.Sprintf("Available for %s", price)

				var isWanted bool
				for i := 0; i < len(wantedCards); i++ {
					if(strings.Contains(name, wantedCards[i])){
						isWanted=true
					}
				}
				if(len(wantedCards)==0||isWanted){
				err := beeep.Alert(name, message, "/icon.png")
				if err != nil {
					panic(err)
				}
				}
			} else {
				fmt.Printf("Not available: %s for %s\n", name, price)
			}
		})

	})

	c.Visit("https://www.steg-electronics.ch/de/search?suche=rtx+4090")
	c.Visit("https://www.steg-electronics.ch/de/search?suche=asus+rtx+3090")

	fmt.Println("!!!!!Smeg crawling finished!!!!!\n")
}
