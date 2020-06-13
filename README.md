# beeper
beeper - play sound via PC speaker

### Installation

    go get github.com/biter777/countries

### Usage
```go
func main() {
	b, _ := beeper.NewBeeper()
	b.Beep(5000, 100*time.Millisecond) // beep sound via PC speaker
	b.BeepItems(beeper.StarWars)       // play StarWars melody via PC speaker :)
	b.Close()
}
```

### Contributing

 Welcome pull requests, bug fixes and issue reports.
 Before proposing a change, please discuss it first by raising an issue.
