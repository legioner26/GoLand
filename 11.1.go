package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Print("Определение количества слов, начинающихся с большой буквы в строке:\n ")
	text := "Пишем текст со словами с Подсчетом Количества начинающихся с Заглавной буквы"
	fmt.Println(text)
	var word string
	var counter int
	for {
		text = strings.Trim(text, " ")
		spaceIndex := strings.Index(text, " ")
		if spaceIndex != -1 {
			word = text[:spaceIndex]
		} else {
			word = text
		}
		if word == strings.Title(word) {
			counter++
		}
		text = strings.TrimLeft(text, word)
		if text == "" {
			fmt.Printf("В тексте %d слов, начинающихся с большой буквы\n", counter)
			break
		}
	}
}
