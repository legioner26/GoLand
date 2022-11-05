package main

import (
	"fmt"
)

func main() {
	fmt.Println("Счастливый билет")
	var one, two, three, four int

	fmt.Println("Введите первую цифру номера билета")
	fmt.Scan(&one)
	fmt.Println("Введите вторую цифру номера билета")
	fmt.Scan(&two)
	fmt.Println("Введите третью цифру номера билета")
	fmt.Scan(&three)
	fmt.Println("Введите четвертую цифру номера билета")
	fmt.Scan(&four)

	if one == four && two == three {
		fmt.Println("зеркальный билет")
	} else if one+two == three+four {
		fmt.Println("счастливый билет")
	} else {
		fmt.Println("обычный билет")
	}

}
