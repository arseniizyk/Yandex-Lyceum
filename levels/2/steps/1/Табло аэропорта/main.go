package main

import (
	"fmt"
)

func PrintFlightRow(flightNumber, cityDepature, cityArrival string, duration float64, placeNumber, gate int, isEnd bool) {
	if !isEnd {
		fmt.Printf("| %s | %s—%s | %d регистрация продолжается |\n", flightNumber, cityDepature, cityArrival, placeNumber)
		return
	}
	fmt.Printf("| %s | %s—%s | регистрация закончилась, проходите к гейту: %d | длительность полёта %.1f часа |\n", flightNumber, cityDepature, cityArrival, gate, duration)

}
