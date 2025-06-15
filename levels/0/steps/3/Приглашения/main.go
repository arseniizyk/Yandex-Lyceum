package main

import "fmt"

func main() {
	var name string
	var room, secret int
	var duration float64

	fmt.Scanln(&name)
	fmt.Scanln(&room)
	fmt.Scanln(&secret)
	fmt.Scanln(&duration)

	result := fmt.Sprintf("Привет, %s! Приглашаю тебя на соревнование по программированию, которое пройдёт, как всегда, в квартире %d. Оно будет длиться примерно %.1f часа. Не забудь секретный пароль для входа: %d.", name, room, duration, secret)

	fmt.Println(result)
}
