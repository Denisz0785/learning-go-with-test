package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {
	t.Run("check print 321Go!", func(t *testing.T) {

		buf := &bytes.Buffer{}
		count := &CountBothSleepAndSpy{}
		want := `3
2
1
Go!`

		CountDown(buf, count)
		got := buf.String()

		if got != want {
			t.Errorf("want %q got %q", want, got)
		}

	})

	t.Run("count calls spy and sleep", func(t *testing.T) {
		c := &CountBothSleepAndSpy{}

		CountDown(c, c)
		want := []string{
			"print",
			"sleep",
			"print",
			"sleep",
			"print",
			"sleep",
			"print",
		}

		if !reflect.DeepEqual(want, c.Calls) {
			t.Errorf("ошибочка, want %v got %v", want, c.Calls)
			t.Log("неувязочка")
		}

	})
}

type SpySleep struct {
	Count int
}

func (s *SpySleep) Sleep() {
	s.Count++
}

type CountBothSleepAndSpy struct {
	Calls []string
}

func (c *CountBothSleepAndSpy) Sleep() {
	c.Calls = append(c.Calls, "sleep")
}

func (c *CountBothSleepAndSpy) Write(p []byte) (n int, e error) {
	c.Calls = append(c.Calls, "print")
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := time.Second * 5

	spyTime := SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("want %v got %v", sleepTime, spyTime.durationSlept)
	}

}
