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

	flag.Parse()

	if *v {
		fmt.Printf("%s %s\n", os.Args[0], getsurei.Version)
	}

	dates := flag.Args()
	if len(dates) < 1 {
		now := time.Now()
		dates = append(dates, now.Format(*f))
	}

	for _, raw := range dates {
		date, err := time.Parse(*f, raw)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to parse date: %s\n", raw)
			continue
		}
		if *a {
			days := getsurei.Getsurei(date)
			fmt.Printf("月齢%.2f\n", days)
		} else {
			fmt.Printf("%s\n", getsurei.Name(date))
		}
	}
}
