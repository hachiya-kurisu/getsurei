// Package getsurei estimates primary moon phases and the age of the moon
package getsurei

import (
	"fmt"
	"math"
	"time"
)

const Version = "0.0.1"

const (
	Shingetsu = iota + 1
	Jougen
	Mangetsu
	Kagen
)

const (
	NewMoon = iota + 1
	FirstQuarter
	FullMoon
	LastQuarter
)

var offsets = map[int]float64{
	Shingetsu: 0,
	Jougen:    7.3826,
	Mangetsu:  14.7652,
	Kagen:     22.1479,
}

const cycle = 29.53059

var reference = time.Date(2019, time.April, 5, 8, 50, 0, 0, time.UTC)

// Getsurei returns the age of the moon in days at time t
func Getsurei(t time.Time) float64 {
	delta := t.Sub(reference).Hours() / 24.0
	mod := math.Mod(delta, cycle)
	if mod < 0 {
		return mod + cycle
	} else {
		return mod
	}
}

// Gessou returns the primary phase of the moon at time t
func Gessou(t time.Time) int {
	g := Getsurei(t)
	switch {
	case g < offsets[Jougen]:
		return Shingetsu
	case g < offsets[Mangetsu]:
		return Jougen
	case g < offsets[Kagen]:
		return Mangetsu
	default:
		return Kagen
	}
}

// Name returns the Japanese name of the primary phase of the moon at time t
func Name(t time.Time) string {
	switch Gessou(t) {
	case Shingetsu:
		return "新月"
	case Jougen:
		return "上限の月"
	case Mangetsu:
		return "満月"
	default:
		return "下弦の月"
	}
}

// Next returns the time of the next phase p after time t
func Next(p int, t time.Time) time.Time {
	g := Getsurei(t)
	ds := fmt.Sprintf("%fh", (cycle-g+offsets[p])*24.0)
	d, _ := time.ParseDuration(ds)
	return t.Add(d)
}
