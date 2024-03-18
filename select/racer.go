package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := MeasureResponseTime(a)
	bDuration := MeasureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func MeasureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
