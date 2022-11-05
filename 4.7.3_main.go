package main

import (
	"fmt"
)

func main() {
	intMax := 100000
	nominal := 100

	fmt.Println("Банкомат.")
	fmt.Println("Введите сумму снятия со счёта:")
	var resultOne int
	fmt.Scan(&resultOne)

	if resultOne%nominal == 0 {
		if resultOne > intMax {
			fmt.Println("Операция не выполнена.")
			fmt.Println("Введите сумму меньше", intMax, "рублей")
		} else {
			fmt.Println("Операция  выполнена.")
			fmt.Println("Вы сняли", resultOne, "рублей.")
		}

	} else {
		fmt.Println("Операция не выполнена. Номинал 100 рублей")
	}

}
