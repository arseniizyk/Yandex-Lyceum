package main

import "fmt"

type Employee struct {
	name     string
	position string
	salary   float64
	bonus    float64
}

func (e Employee) CalculateTotalSalary() {
	total := e.salary + e.bonus
	fmt.Println("Employee:", e.name)
	fmt.Println("Position:", e.position)
	fmt.Printf("Total Salary: %.2f\n", total)
}
