package crawler

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func getSleepSeconds() int {
	sleepSecondsRaw := os.Getenv("SLEEP_SECONDS")
	sleepSeconds, err := strconv.Atoi(sleepSecondsRaw)
	if err != nil {
		defaultSleepSeconds := 60
		fmt.Printf("SLEEP_SECONDS environment variable is not a valid integer, using a default value of %d seconds.", sleepSeconds)
		return defaultSleepSeconds
	}
	return sleepSeconds
}

func TakeABreak() {
	sleepSeconds := getSleepSeconds()
	fmt.Printf("Taking a break for %d secs...", sleepSeconds)
	time.Sleep(time.Duration(sleepSeconds) * time.Second)
}
