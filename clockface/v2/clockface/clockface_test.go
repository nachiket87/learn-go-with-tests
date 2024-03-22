package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadian(t *testing.T) {
	thirtySeconds := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)
	want := math.Pi
	got := secondsInRadians(thirtySeconds)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

}
