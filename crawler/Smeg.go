package crawler

import (
	"fmt"
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
				unwantedCards:= []string{"Zotac", "PNY", "KFA2", "Manli", "Inno3D", "Palit", "Gainward"}
				var isUnwanted bool
				for i := 0; i < len(unwantedCards); i++ {
					if(strings.Contains(name, unwantedCards[i])){
						isUnwanted=true
					}
				}
				if(!isUnwanted){
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

	// c.OnHTML("strong.sc-1ttlt4t-1.kRYPrq", func(e *colly.HTMLElement) {
	// 	price = e.Text
	// })

	// c.OnHTML("button.sc-1olg58b-0.cXWAZg.sc-185g6wq-6.bvMIBV", func(e *colly.HTMLElement) {
	// 	_, disabled = e.DOM.Attr("disabled")
	// 	if disabled {
	// 		fmt.Printf("Not available: %s for %s\n", name, price)
	// 	} else {
	// 		fmt.Printf("Available: %s for %s\n", name, price)
	// 	}
	// })

	c.Visit("https://www.steg-electronics.ch/de/search?suche=rtx+4090")
	c.Visit("https://www.steg-electronics.ch/de/search?suche=rtx%204090&p=2")

	fmt.Println("!!!!!Smeg crawling finished!!!!!\n")
}
