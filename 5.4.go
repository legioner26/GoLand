package main

import (
	"fmt"
)

func main() {
	var price, oneNominal, twoNominal, threeNominal int

	fmt.Println("Введите стоимость товара")
	fmt.Scan(&price)
	fmt.Println("Введите номинал первой монеты")
	fmt.Scan(&oneNominal)
	fmt.Println("Введите номинал второй монеты")
	fmt.Scan(&twoNominal)
	fmt.Println("Введите номинал третьей монеты")
	fmt.Scan(&threeNominal)

	summ := oneNominal + twoNominal + threeNominal

	if (price == oneNominal || price == twoNominal || price == threeNominal) || (price == summ) || (price == oneNominal+twoNominal || price == oneNominal+threeNominal || price == twoNominal+threeNominal) {
		fmt.Println("Можно расплатиться без сдачи")
	} else {
		fmt.Println("Без сдачи не расплатиться")
	}
}
