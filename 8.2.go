package main

import (
	"fmt"
)

func main() {
	fmt.Println("Дни недели.")
	fmt.Println("Введите будний день недели - пн, вт, ср, чт, пт:")

	var day string
	_, _ = fmt.Scan(&day)

	switch day {
	case "пн":
		fmt.Println("понедельник")
		fallthrough
	case "вт":
		fmt.Println("вторник")
		fallthrough
	case "ср":
		fmt.Println("среда")
		fallthrough
	case "чт":
		fmt.Println("четверг")
		fallthrough
	case "пт":
		fmt.Println("пятница")
	default:
		fmt.Println("Не известен день недели")
	}

}
