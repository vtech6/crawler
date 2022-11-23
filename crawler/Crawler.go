package crawler

import (
	_ "github.com/joho/godotenv/autoload"
)

func Crawler() {
	Asus()
	Smeg()
	Digitec()

	TakeABreak()

	Crawler()
}
