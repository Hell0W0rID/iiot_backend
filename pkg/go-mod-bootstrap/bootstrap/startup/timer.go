/*******************************************************************************
 *******************************************************************************/

package startup

import (
	"time"

	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/environment"
)

// Timer contains references to dependencies required by the startup timer implementation.
type Timer struct {
	startTime time.Time
	duration  time.Duration
	interval  time.Duration
}

// NewStartUpTimer is a factory method that returns an initialized Timer receiver struct.
func NewStartUpTimer(serviceKey string) Timer {
	startup := environment.GetStartupInfo(serviceKey)

	return Timer{
		startTime: time.Now(),
		duration:  time.Second * time.Duration(startup.Duration),
		interval:  time.Second * time.Duration(startup.Interval),
	}
}

// NewTimer is a factory method that returns a Timer initialized with passed in duration and interval.
func NewTimer(duration int, interval int) Timer {
	return Timer{
		startTime: time.Now(),
		duration:  time.Second * time.Duration(duration),
		interval:  time.Second * time.Duration(interval),
	}
}

// SinceAsString returns the time since the timer was created as a string.
func (t Timer) SinceAsString() string {
	return time.Since(t.startTime).String()
}

// RemainingAsString returns the time remaining on the timer as a string.
func (t Timer) RemainingAsString() string {

	remaining := t.duration - time.Since(t.startTime)
	if remaining < 0 {
		remaining = 0
	}
	return remaining.String()
}

// HasNotElapsed returns whether or not the duration specified during construction has elapsed.
func (t Timer) HasNotElapsed() bool {
	return time.Now().Before(t.startTime.Add(t.duration))
}

// SleepForInterval pauses execution for the interval specified during construction.
func (t Timer) SleepForInterval() {
	time.Sleep(t.interval)
}
