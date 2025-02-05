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

		if *g {
			fmt.Printf("%s\n", getsurei.Name(date))
		}

		if *a {
			days := getsurei.Getsurei(date)
			fmt.Printf("月齢%.2f\n", days)
		}

		if *s {
			shingetsu := getsurei.Next(getsurei.Shingetsu, date)
			fmt.Printf("次の新月：%s\n", shingetsu.Format(*f))
		}

		if *j {
			jougen := getsurei.Next(getsurei.Jougen, date)
			fmt.Printf("次の上弦の月：%s\n", jougen.Format(*f))
		}

		if *m {
			mangetsu := getsurei.Next(getsurei.Mangetsu, date)
			fmt.Printf("次の満月：%s\n", mangetsu.Format(*f))
		}

		if *k {
			kagen := getsurei.Next(getsurei.Kagen, date)
			fmt.Printf("次の下弦の月：%s\n", kagen.Format(*f))
		}
	}
}
