package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3
const countdownEnd = 0
const sleepSec = 1

func main() {
	sleeper := &ConfigurableSleeper{sleepSec * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > countdownEnd; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprintln(writer, finalWord)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type Sleeper interface {
	Sleep()
}
