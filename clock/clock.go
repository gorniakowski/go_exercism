package clock

import "fmt"

//Clock represneting time
type Clock struct {
	hours   int
	minutes int
}

func (time Clock) String() string {
	return fmt.Sprintf("%02d:%02d", time.hours, time.minutes)
}

//Add given minutes to a clock
func (time Clock) Add(mins int) Clock {
	return New(time.hours, time.minutes+mins)
}

//Subtract given numbers form clock
func (time Clock) Subtract(mins int) Clock {
	return New(time.hours, time.minutes-mins)
}

//New Clock constructor
func New(hour, minute int) Clock {

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
