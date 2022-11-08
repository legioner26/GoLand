package main

import (
	"fmt"
)

func main() {
	fmt.Println("Времена года")
	fmt.Println("Введите месяц")

	var month string
	_, _ = fmt.Scan(&month)

	switch month {
	case "декабрь", "январь", "февраль":
		fmt.Println("Время года - зима")
	case "март", "апрель", "май":
		fmt.Println("Время года - весна")
	case "июнь", "июль", "август":
		fmt.Println("Время года - лето")
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println("Время года - осень")
	default:
		fmt.Println("Не известен месяц !!!")
	}
}
