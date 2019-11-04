package clock

import "fmt"

type Clock struct {
	hours   int
	minutes int
}

func (time Clock) String() string {
	return fmt.Sprintf("%02d:%02d", time.hours, time.minutes)
}

func New(hour, minute int) Clock {

	if hour == 24 { //we don't want 24 but 00
		hour = 0
	}
	for minute >= 60 {
		hour++
		minute -= 60
	}
	for minute < 0 {
		hour--
		minute += 60
	}

	for hour < 0 {
		hour += 24
	}

	return Clock{hour % 24, minute}
}
