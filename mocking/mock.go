package mocking

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

func CountDown(w io.Writer, sleep Sleeper) {

	for i := 3; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleep.Sleep()
	}

	fmt.Fprint(w, "Go!")
}

type ConfigurableSleeper struct {
	Duration time.Duration
	Sleepf   func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.Sleepf(c.Duration)
}
