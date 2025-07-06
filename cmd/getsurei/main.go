package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"blekksprut.net/getsurei"
)

func main() {
	a := flag.Bool("a", false, "print the age of the moon in days")
	v := flag.Bool("v", false, "version")
	f := flag.String("f", "2006.01.02", "date format")

	s := flag.Bool("s", false, "print next new moon")
	j := flag.Bool("j", false, "print next first quarter moon")
	m := flag.Bool("m", false, "print next full moon")
	k := flag.Bool("k", false, "print next last quarter moon")

	g := flag.Bool("g", true, "print corresponding moon phase")

	l := flag.String("l", "jp", "language (ja/en/no)")
	
	tz := flag.String("tz", "", "time zone")

	flag.Parse()

	var loc *time.Location
	var err error
	if *tz != "" {
		loc, err = time.LoadLocation(*tz)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid time zone: %s\n", *tz)
			os.Exit(1)
		}
	} else {
		loc = time.Local
	}

	if *v {
		fmt.Printf("%s %s\n", os.Args[0], getsurei.Version)
		os.Exit(0)
	}

	dates := flag.Args()
	if len(dates) < 1 {
		now := time.Now().In(loc)
		dates = append(dates, now.Format(*f))
	}

	for _, raw := range dates {
		date, err := time.Parse(*f, raw)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to parse date: %s\n", raw)
			continue
		}

		if *g {
			fmt.Printf("%s\n", getsurei.Name(date, *l))
		}

		if *a {
			days := getsurei.Getsurei(date)
			fmt.Printf("%s%.2f\n", getsurei.T("getsurei", *l), days)
		}

		if *s {
			shingetsu := getsurei.Next(getsurei.Shingetsu, date).In(loc)
			fmt.Printf("%s：%s\n", getsurei.T("shingetsu", *l), shingetsu.Format(*f))
		}

		if *j {
			jougen := getsurei.Next(getsurei.Jougen, date).In(loc)
			fmt.Printf("%s：%s\n", getsurei.T("jougen", *l), jougen.Format(*f))
		}

		if *m {
			mangetsu := getsurei.Next(getsurei.Mangetsu, date).In(loc)
			fmt.Printf("%s：%s\n", getsurei.T("mangetsu", *l), mangetsu.Format(*f))
		}

		if *k {
			kagen := getsurei.Next(getsurei.Kagen, date).In(loc)
			fmt.Printf("%s：%s\n", getsurei.T("kagen", *l), kagen.Format(*f))
		}
	}
}
