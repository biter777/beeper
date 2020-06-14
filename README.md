# beeper
beeper - play sound via PC speaker<br/><br/>
[![GoDoc](https://godoc.org/github.com/biter777/beeper?status.svg)](https://godoc.org/github.com/biter777/beeper)
[![GoDev](https://img.shields.io/badge/godev-reference-5b77b3)](https://pkg.go.dev/github.com/biter777/beeper?tab=doc)
[![Go Walker](https://img.shields.io/badge/gowalker-reference-5b77b3)](https://gowalker.org/github.com/biter777/beeper)
[![GolangCI](https://golangci.com/badges/github.com/biter777/beeper.svg?style=flat)](https://golangci.com/r/github.com/biter777/beeper)
[![GoReport](https://goreportcard.com/badge/github.com/biter777/beeper)](https://goreportcard.com/report/github.com/biter777/beeper)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/08eb1d2ff62e465091b3a288ae078a96)](https://www.codacy.com/manual/biter777/beeper?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=biter777/beeper&amp;utm_campaign=Badge_Grade)
[![License](https://img.shields.io/badge/License-BSD%202--Clause-brightgreen.svg)](https://opensource.org/licenses/BSD-2-Clause)
[![Build status](https://ci.appveyor.com/api/projects/status/t9lpor9o8tpacpmr/branch/master?svg=true)](https://ci.appveyor.com/project/biter777/beeper/branch/master)
[![Build Status](https://github.com/go-vgo/robotgo/workflows/Go/badge.svg)](https://github.com/go-vgo/robotgo/commits/master)
[![Dependencies Free](https://img.shields.io/badge/dependencies-free-brightgreen)](https://pkg.go.dev/github.com/biter777/beeper?tab=imports)
[![Gluten Free](https://img.shields.io/badge/gluten-free-brightgreen)](https://www.scsglobalservices.com/services/gluten-free-certification)
[![DepShield Badge](https://depshield.sonatype.org/badges/biter777/beeper/depshield.svg)](https://depshield.github.io)
[![Stars](https://img.shields.io/github/stars/biter777/beeper?label=Please%20like%20us&style=social)](https://github.com/biter777/beeper/stargazers)

### Installation

    go get github.com/biter777/beeper

### Usage
```go
func main() {
	b, _ := beeper.NewBeeper()
	b.Beep(5000, 100*time.Millisecond) // beep sound via PC speaker, 5000 hz, 100 ms
	b.BeepItems(beeper.StarWars())       // play StarWars melody via PC speaker :)
	b.Close()
}
```

### Contributing

 Welcome pull requests, bug fixes and issue reports.<br/>
 Before proposing a change, please discuss it first by raising an issue.<br/>
 <b>Star us</b>. Give us a star, please, if it's not against your religion :)
