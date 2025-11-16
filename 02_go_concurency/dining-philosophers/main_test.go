package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("incorrect length of slice; expected 5 but got %d", len(orderFinished))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	tests := []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Second * 0},
		{"quater second delay", time.Millisecond * 25},
		{"helf second delay", time.Millisecond * 500},
	}

	for _, test := range tests {
		orderFinished = []string{}

		eatTime = test.delay
		sleepTime = test.delay
		thinkTime = test.delay

		dine()
		if len(orderFinished) != 5 {
			t.Errorf("%s: incorrect length of slice; expected 5 but got %d", test.name, len(orderFinished))
		}
	}

}
