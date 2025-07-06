// Package getsurei estimates primary moon phases and the age of the moon
package getsurei

import (
	"math"
	"time"
)

const Version = "0.0.4"

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

var locales = map[string]map[string]string{
	"jp": {
		"getsurei":  "月齢",
		"shingetsu": "新月",
		"jougen":    "上弦の月",
		"mangetsu":  "満月",
		"kagen":     "下弦の月",
	},
	"en": {
		"getsurei":  "moon age",
		"shingetsu": "new moon",
		"jougen":    "first quarter",
		"mangetsu":  "full moon",
		"kagen":     "third quarter",
	},
	"no": {
		"getsurei":  "månealder",
		"shingetsu": "nymåne",
		"jougen":    "første kvarter",
		"mangetsu":  "fullmåne",
		"kagen":     "siste kvarter",
	},
}

func T(key, locale string) string {
	translation, ok := locales[locale]
	if ok {
		translated, ok := translation[key]
		if ok {
			return translated
		}
	}
	return locales["jp"][key]
}

var offsets = map[int]float64{
	Shingetsu: 0,
	Jougen:    7.3826,
	Mangetsu:  14.7652,
	Kagen:     22.1479,
}

const cycle = 29.530589

var reference = time.Date(2019, time.April, 5, 8, 50, 0, 0, time.UTC)

// Getsurei returns the age of the moon in days at time t
func Getsurei(t time.Time) float64 {
	delta := t.Sub(reference).Hours() / 24.0
	mod := math.Mod(delta, cycle)
	if mod < 0 {
		return mod + cycle
	}
	return mod
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

// Name returns the localized name of the primary phase of the moon at time t
func Name(t time.Time, locale string) string {
	switch Gessou(t) {
	case Shingetsu:
		return T("shingetsu", locale)
	case Jougen:
		return T("jougen", locale)
	case Mangetsu:
		return T("mangetsu", locale)
	default:
		return T("kagen", locale)
	}
}

// Next returns the time of the next phase p after time t
func Next(p int, t time.Time) time.Time {
	g := Getsurei(t)
	days := offsets[p] - g
	if days < 0 {
		days += cycle
	}
	hours := days * 24
	return t.Add(time.Duration(hours * float64(time.Hour)))
}
