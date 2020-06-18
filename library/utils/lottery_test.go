package utils

import (
	"fmt"
	"testing"
	"time"
)

// TestLottery test lottery
func TestLottery(t *testing.T) {
	start := time.Now()
	_, err := Lottery(0.0050)
	if err != nil {
		t.Fail()
	}

	end := time.Now()

	fmt.Printf("time: %v\n", end.Sub(start).Seconds())
}
