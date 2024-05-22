package selectconcur

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = time.Second * 10

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("error of waiting response more than 10 second for %s and %s", a, b)
	}
}

func ping(s string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(s)
		close(ch)
	}()
	return ch
}

func measureResponseTime(a string) time.Duration {
	start := time.Now()
	http.Get(a)
	return time.Since(start)
}
