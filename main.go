package main

import (
	"lgwt/mocking"
	"os"
	"time"
)

func main() {
	sleepTime := 1 * time.Second
	sleepFunc := time.Sleep
	sleeper := &mocking.ConfigurableSleeper{}
	sleeper.UpdateConfig(sleepTime, sleepFunc)
	mocking.Countdown(os.Stdout, sleeper)
}
