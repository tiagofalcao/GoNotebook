// Package clock implements only minute and hours clock.
package clock

import "fmt"

// Clock is a clock that store only hours and minutes.
type Clock int

// day is the amount of minutes in one day.
const day = 24 * 60

// New returns a new Clock with h hours and m minutes.
func New(h, m int) Clock {
	return Clock(0).Add(h*60 + m)
}

// Add returns a new clock with m minutes in addition to clock c.
func (c Clock) Add(m int) Clock {
	c += Clock(m)
	c %= day
	if c < 0 {
		c += day
	}
	return c
}

// Subtract returns a new clock with m minutes in subtraction to clock c.
func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

// String returns the clock as string with two digits to hours and minutes.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}
