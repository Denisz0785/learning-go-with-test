package selectconcur

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration > bDuration {
		return b
	}
	return a
}

func measureResponseTime(a string) time.Duration {
	start := time.Now()
	http.Get(a)
	return time.Since(start)
}
