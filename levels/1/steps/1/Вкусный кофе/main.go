package main

import "fmt"

func main() {
	var rubles, pennies, price int

	fmt.Scan(&rubles, &pennies, &price)

	if pennies >= 100 {
		rubles += pennies / 100
	}

	if rubles < price {
		fmt.Println("Стоит подкопить")
		return
	}

	fmt.Println("Сегодня будет вкусный кофе!")
}
