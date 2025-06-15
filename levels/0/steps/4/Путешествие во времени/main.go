package main

import (
	"fmt"
	"time"
)

func main() {
	var futureYear, pastYear string

	fmt.Scan(&futureYear, &pastYear)

	const hoursInYear = 8760
	futureYearTime, err := time.Parse("2006-01-02", futureYear)
	if err != nil {
		panic(err)
	}
	pastYearTime, err := time.Parse("2006-01-02", pastYear)
	if err != nil {
		panic(err)
	}

	dur := futureYearTime.Year() - pastYearTime.Year()
	fmt.Println(dur, "year ago")
}
