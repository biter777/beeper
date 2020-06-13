package beeper

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"
)

// linuxBeeper - PC-speaker beeper for Linux
type linuxBeeper struct {
	spkr  *os.File
	evdev bool
}

const (
	eventType = 0x12    // linux/input-event-codes.h
	eventCode = 0x02    // linux/input-event-codes.h
	kiocSound = 0x4B2F  // linux/kd.h, 0x4B2F = sound generation start; 0 = stop generation
	clockRate = 1193180 // 1.193180 MHz = frequency of the original PC XT, lol :)
)

// helpInfo - help information how to enable a PC-speaker on Linux
const helpInfo = "1) addgroup --system beep\n" +
	"2) usermod -a -G beep USERNAME\n" +
	"3) Add udev rule in /usr/lib/udev/rules.d/70-pcspkr-beep.rules or /lib/udev/rules.d/70-pcspkr-beep.rules (the exact location depends on your distribution).\n" +
	"# Give write access to the PC speaker only to the \"beep\" group\n" +
	"ACTION==\"add\", SUBSYSTEM==\"input\", ATTRS{name}==\"PC Speaker\", ENV{DEVNAME}!=\"\", GROUP=\"beep\", MODE=\"0620\"\n" +
	"4) reboot\n" +
	"5) modprobe -r pcspkr; sleep 3; modprobe pcspkr\n" +
	"Maybe you must edit /etc/modprobe.d/blacklist.conf for enable pcspkr (or snd_pcsp; use \"lsmod | grep pcsp\" for help)."

// event - linux/input.h event struct
type event struct {
	Time  syscall.Timeval // time in seconds since epoch at which event occurred
	Type  uint16          // event type
	Code  uint16          // event code related to the event type
	Value int32           // event value related to the event type
}

func newEvent(freq int32) *event {
	return &event{
		Type:  eventType,
		Code:  eventCode,
		Value: freq,
	}
}

// NewBeeper - NewBeeper
func NewBeeper() (Beeper, error) {
	return newLinuxBeeper()
}

// newLinuxBeeper - newLinuxBeeper
func newLinuxBeeper() (b *linuxBeeper, err error) {
	b = &linuxBeeper{}
	b.spkr, err = os.OpenFile("/dev/tty0", os.O_WRONLY, 0644)
	if err != nil {
		var err2 error
		b.spkr, err2 = os.OpenFile("/dev/input/by-path/platform-pcspkr-event-spkr", os.O_WRONLY, 0644)
		if err2 != nil {
			return nil, fmt.Errorf("newLinuxBeeper:: OpenFile error: %v %w\n%v", err, err2, helpInfo)
		}
		b.evdev = true
	}
	return b, nil
}

// Close - Close
func (b *linuxBeeper) Close() error {
	return b.spkr.Close()
}

// Beep - Beep
// @freq = frequency
// @duration = duration :)
func (b *linuxBeeper) Beep(freq int, duration time.Duration) error {
	if b.evdev {
		e := newEvent(int32(freq))
		data := *(*[unsafe.Sizeof(*e)]byte)(unsafe.Pointer(e))
		// Start sound
		_, err := b.spkr.Write(data[:])
		if err != nil {
			return err
		}
		time.Sleep(duration)
		e.Value = 0
		data = *(*[unsafe.Sizeof(*e)]byte)(unsafe.Pointer(e))
		// Stop sound
		_, err = b.spkr.Write(data[:])
		return err
	}

	// Start sound
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, b.spkr.Fd(), kiocSound, uintptr(clockRate/freq))
	if errno != 0 {
		return errors.New(errno.Error())
	}

	time.Sleep(duration)
	// Stop sound
	_, _, errno = syscall.Syscall(syscall.SYS_IOCTL, b.spkr.Fd(), kiocSound, 0)
	if errno != 0 {
		return errors.New(errno.Error())
	}

	return nil
}

// BeepItem - BeepItem
func (b *linuxBeeper) BeepItem(item *BeepItem) error {
	return b.Beep(item.Freq, item.Duration)
}

// BeepItems - BeepItems
func (b *linuxBeeper) BeepItems(items []BeepItem) (err error) {
	for _, item := range items {
		err = b.Beep(item.Freq, item.Duration)
		if err != nil {
			return err
		}
		time.Sleep(item.Pause)
	}
	return nil
}
