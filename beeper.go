// Package beeper - play sound via PC speaker
package beeper

// Usage:
//  b, _ := beeper.NewBeeper()
// 	b.Beep(5000, 100*time.Millisecond) // beep sound via PC speaker, 5000 hz, 100 ms
// 	b.BeepItems(beeper.StarWars())       // play StarWars melody via PC speaker :)
// 	b.Close()

import (
	"os"
	"time"
)

// Beeper - PC-speaker beeper
type Beeper interface {
	Beep(freq int, duration time.Duration) error
	BeepItem(item *BeepItem) error
	BeepItems(items []BeepItem, speed float64) error
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
