package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Исходная строка:")
	text := "a10 10 20b 20 30c30 30 dd"
	fmt.Println(text)

	var word string
	var result string

	for {
		text = strings.Trim(text, " ")
		spaceIndex := strings.Index(text, " ")

		if spaceIndex > 0 {
			word = text[:spaceIndex]
		} else {
			word = text
		}
		num, err := strconv.ParseInt(word, 10, 0)
		if err != nil {

			text = strings.Trim(strings.Replace(text, word, "", 1), " ")

		} else {
			result += " " + strconv.Itoa(int(num))
			text = strings.Trim(strings.Replace(text, word, "", 1), " ")

		}
		if text == "" {
			result = strings.Trim(result, " ")
			fmt.Println("В строке содержатся числа в десятичном формате:")
			fmt.Println(result)
			break
		}
	}
}
