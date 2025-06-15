package main

import (
	"fmt"
)

type Animal interface {
	MakeSound()
}

type Dog struct{}
type Cat struct{}

func (c Cat) MakeSound() {
	fmt.Println("Мяу!")
}

func (d Dog) MakeSound() {
	fmt.Println("Гав!")
}
