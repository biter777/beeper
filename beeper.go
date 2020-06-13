package beeper

import (
	"os"
	"time"
)

// Beeper - PC-speaker beeper
type Beeper interface {
	Beep(freq int, duration time.Duration) error
	BeepItem(item *BeepItem) error
	BeepItems(items []BeepItem) error
	Close() error
}

// BeepItem - Beep Item (freq and duration)
type BeepItem struct {
	Freq     int
	Duration time.Duration
	Pause    time.Duration // Pause after sound off, used in BeepItems()
}

// SimpleBeep - just a simple beep via stdout
func SimpleBeep() error {
	_, err := os.Stdout.Write([]byte{7})
	return err
}
