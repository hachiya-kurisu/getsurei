package getsurei_test

import (
	"math"
	"testing"
	"time"

	"blekksprut.net/getsurei"
)

var layout = "2006.01.02"

func TestGetsurei(t *testing.T) {
	date, _ := time.Parse(layout, "2025.01.27")
	getsurei := getsurei.Getsurei(date)
	if math.Round(getsurei) != 27 {
		t.Errorf("oops: %f", getsurei)
	}
}

func TestShingetsu(t *testing.T) {
	date, _ := time.Parse(layout, "2025.01.30")
	gessou := getsurei.Name(date, "jp")
	if gessou != "新月" {
		t.Errorf("oops: %s", gessou)
	}
}

func TestJougen(t *testing.T) {
	date, _ := time.Parse(layout, "2025.01.14")
	gessou := getsurei.Name(date, "jp")
	if gessou != "上弦の月" {
		t.Errorf("oops: %s", gessou)
	}
}

func TestMangetsu(t *testing.T) {
	date, _ := time.Parse(layout, "2025.01.20")
	gessou := getsurei.Name(date, "jp")
	if gessou != "満月" {
		t.Errorf("oops: %s", gessou)
	}
}

func TestKagen(t *testing.T) {
	date, _ := time.Parse(layout, "2025.01.27")
	gessou := getsurei.Name(date, "jp")
	if gessou != "下弦の月" {
		t.Errorf("oops: %s", gessou)
	}
}

func TestBeforeReferenceDate(t *testing.T) {
	date, _ := time.Parse(layout, "2019.01.27")
	gessou := getsurei.Name(date, "jp")
	if gessou != "満月" {
		t.Errorf("oops: %s", gessou)
	}
}

func TestLocaleFallback(t *testing.T) {
	date, _ := time.Parse(layout, "2019.01.27")
	gessou := getsurei.Name(date, "n/a")
	if gessou != "満月" {
		t.Errorf("oops: %s", gessou)
	}
}


func TestNextShingetsu(t *testing.T) {
	date, _ := time.Parse(layout, "2025.01.27")
	shingetsu := getsurei.Next(getsurei.Shingetsu, date)
	formatted := shingetsu.Format("2006.01.02 15:04")
	if formatted != "2025.01.29 13:41" {
		t.Errorf("oops: %s", formatted)
	}
}

func TestNextMangetsu(t *testing.T) {
	date, _ := time.Parse(layout, "2025.01.27")
	shingetsu := getsurei.Next(getsurei.Mangetsu, date)
	formatted := shingetsu.Format("2006.01.02 15:04")
	if formatted != "2025.02.13 08:03" {
		t.Errorf("oops: %s", formatted)
	}
}
