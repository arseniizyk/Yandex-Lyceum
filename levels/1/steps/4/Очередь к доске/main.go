package main

import (
	"fmt"
)

func main() {
	students := make([]string, 5)
	studentsInQueue := 0
	for {
		var input string
		var place int
		fmt.Scan(&input)
		switch input {
		case "очередь":
			printQueue(students)
			continue
		case "конец":
			printQueue(students)
			return
		case "количество":
			printRemainders(studentsInQueue)
			continue
		default:
			fmt.Scan(&place)
			err := handlePlace(place, studentsInQueue, students)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			students[place-1] = input
			studentsInQueue++
		}
	}
}

func handlePlace(place, studentsInQueue int, students []string) error {
	if place <= 0 || place > 5 {
		return fmt.Errorf("Запись на место номер %d невозможна: некорректный ввод", place)
	} else if studentsInQueue == 5 {
		return fmt.Errorf("Запись на место номер %d невозможна: очередь переполнена", place)
	} else if students[place-1] != "" {
		return fmt.Errorf("Запись на место номер %d невозможна: место уже занято", place)
	}
	return nil
}

func printRemainders(studentsInQueue int) {
	fmt.Println("Осталось свободных мест:", 5-studentsInQueue)
	fmt.Println("Всего человек в очереди:", studentsInQueue)
}

func printQueue(slice []string) {
	for i, item := range slice {
		if item == "" {
			fmt.Printf("%d. -\n", i+1)
			continue
		}
		fmt.Printf("%d. %s\n", i+1, item)
	}
}
