package clock

import "fmt"

// Clock a new type to model a clock with just hours and minutes
type Clock struct {
	hour   int
	minute int
}

// MINUTES is number of minutes per hour
const MINUTES = 60

func convertHelper(hour, minute int) (int, int) {
	totalMinutes := hour*MINUTES + minute
	for totalMinutes < 0 {
		totalMinutes = 24*MINUTES + totalMinutes
	}
	adjustedMinute := totalMinutes % MINUTES
	adjustedHour := ((totalMinutes - adjustedMinute) / MINUTES)
	adjustedHour = adjustedHour % 24
	return adjustedHour, adjustedMinute
}

// New constructs a pointer to clock struct
func New(hour, minute int) Clock {
	adjustedHour, adjustedMinute := convertHelper(hour, minute)
	c := Clock{adjustedHour, adjustedMinute}
	return c
}

// Clock.String returns string representation of time
func (clock Clock) String() string {
	return fmt.Sprintf("%.2d:%.2d", clock.hour, clock.minute)
}

// Add constructs a new clock struct after addition/subtration of minutes
func (clock Clock) Add(minute int) Clock {
	return New(clock.hour, clock.minute+minute)
}
