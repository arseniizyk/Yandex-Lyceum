package main

import "fmt"

func PrettyArrayOutput(array [9]string) {
	for i, text := range array {
		if i < 7 {
			fmt.Printf("%d я уже сделал: %s\n", i+1, text)
			continue
		}
		fmt.Printf("%d не успел сделать: %s\n", i+1, text)
	}

}
