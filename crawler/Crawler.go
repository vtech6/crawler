package crawler

import (
	"fmt"
	"time"
)

func Crawler() {
	Smeg()
	Digitec()
	fmt.Println("Taking a break for a minute.")
	time.Sleep(1 * time.Minute)
	Crawler()
}
