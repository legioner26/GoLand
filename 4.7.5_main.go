package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите день недели:")
	var dayNumber int
	fmt.Scan(&dayNumber)
	fmt.Println("Введите число гостей:")
	var countPersons int
	fmt.Scan(&countPersons)
	fmt.Println("Введите сумму чека:")
	var checkSum int
	fmt.Scan(&checkSum)
	var discount, service, total int

	if dayNumber == 5 {
		discount = checkSum / 20
		fmt.Println("Скидка по пятницам:", discount)

	}

	if countPersons > 5 {
		service += checkSum / 10
		fmt.Println("Надбавка за обслуживание 10%:", service)

	}

	total = checkSum - discount + service
	fmt.Println("Сумма к оплате:", total)
}
