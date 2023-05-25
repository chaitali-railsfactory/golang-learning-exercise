// Implement a clock that handles times without dates.

// You should be able to add and subtract minutes to it.

// Two clocks that represent the same time should be equal to each other.
package main

import (
	"fmt"
)

// Define the Clock type here.
type Clock struct {
	minutes int
	hours   int
}

func New(h, m int) Clock {
	c := Clock{
		minutes: m,
		hours:   h,
	}
	return c.FormatTime()
}

func (c Clock) Add(m int) Clock {
	return New(c.hours, c.minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hours, c.minutes-m)
}

func (c *Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

func (c Clock) FormatTime() Clock {
	for c.minutes > 59 {
		c.minutes = c.minutes - 60
		c.hours += 1
	}
	for c.minutes < 0 {
		c.minutes = c.minutes + 60
		c.hours -= 1
	}
	for c.hours > 23 {
		c.hours = c.hours - 24
	}
	for c.hours < 0 {
		c.hours = 24 + c.hours
	}
	return c
}
func main() {
	var clock, subTime, addTime Clock
	clock = New(72, 8640)
	fmt.Printf("Time is hours:%d, min:%d \n", clock.hours, clock.minutes)
	subTime = New(2, 20).Subtract(3000)
	fmt.Printf("Subtracted time output is hours:%d, min:%d \n", subTime.hours, subTime.minutes)
	addTime = New(1, 1).Add(3500)
	fmt.Printf("Subtracted time output is hours:%d, min:%d \n", addTime.hours, addTime.minutes)
}
