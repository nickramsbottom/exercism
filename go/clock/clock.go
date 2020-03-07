package clock

import "fmt"

// Clock has functionality for displaying and manipulating time
type Clock struct {
	minutes int
}

// New creates a new instance of Clock
func New(h, m int) Clock {
	minutes := (h*60 + m) % (60 * 24)
	for minutes < 0 {
		minutes += 60 * 24
	}
	return Clock{minutes}
}

// String prints the current time as a string
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

// Add creates a new Clock instance with the time increased
// by a given number of minutes
func (c Clock) Add(minutes int) Clock {
	c.minutes += minutes
	return New(0, c.minutes)
}

// Subtract creates a new Clock instance with the time decreased
// by a given number of minutes
func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}
