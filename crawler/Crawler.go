package crawler

import (
	_ "github.com/joho/godotenv/autoload"
)

func Crawler() {
	Smeg()
	Digitec()

	TakeABreak()

	Crawler()
}
