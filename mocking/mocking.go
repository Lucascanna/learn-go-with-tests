package mocking

import (
	"fmt"
	"io"
)

// Sleeper interface
type Sleeper interface {
	Sleep()
}

// Countdown func
func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	sleeper.Sleep()
	fmt.Fprint(writer, "Go!")
}
