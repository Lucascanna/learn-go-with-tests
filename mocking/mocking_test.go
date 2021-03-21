package mocking

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}

	wantSleepCalls := 4
	gotSleepCalls := spySleeper.Calls
	if gotSleepCalls != wantSleepCalls {
		t.Errorf("got %d want %d given", gotSleepCalls, wantSleepCalls)
	}
}
