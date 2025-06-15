package main

import (
	"fmt"
	"time"
)

const coinsInRuble = 100

func main() {
	var startWork, name, lastName, fathersName string
	var payAmountFirst, payAmountSecond, payAmountThird float64

	fmt.Scan(&startWork, &name, &lastName, &fathersName, &payAmountFirst, &payAmountSecond, &payAmountThird)

	startWorkTime, err := time.Parse("02.01.2006", startWork)
	if err != nil {
		panic(err)
	}

	signDate := startWorkTime.AddDate(0, 0, 15).Format("02.01.2006")

	totalAmount := payAmountFirst + payAmountSecond + payAmountThird
	payAmountRubles := int(totalAmount)
	payAmountCoins := int((totalAmount - float64(payAmountRubles)) * coinsInRuble)
	result := fmt.Sprintf(`Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.
Дата подписания договора: %v. Просим вас подойти в офис в любое удобное для вас время в этот день.
Общая сумма выплат составит %d руб. %d коп.

С уважением,
Гл. бух. Иванов А.Е.`, lastName, name, fathersName, signDate, payAmountRubles, payAmountCoins)

	fmt.Println(result)
}
