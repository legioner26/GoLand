package main

import (
	"fmt"
	"strings"
)

const lenSentences = 4
const lenChars = 5

func main() {
	sentences := [lenSentences]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}

	fmt.Println(sentences)

	chars := [lenChars]rune{'H', 'E', 'L', 'П', 'М'}
	parseTest(sentences, chars)
}

func explode(sentences [lenSentences]string) (word [lenSentences]string) {
	for i := 0; i < lenSentences; i++ {
		word[i] = strings.ToUpper(strings.Fields(sentences[i])[1])
		//fmt.Println(strings.Fields(sentences[i])[1])
	}

	return
}

func parseTest(sentences [lenSentences]string, chars [lenChars]rune) {
	word := explode(sentences)

	index := -1
	for _, val := range word {
		for _, run := range chars {
			index = strings.IndexRune(val, run)
			if index >= 0 {
				fmt.Println(string(run), "индекс", index)
			}
		}
	}

}
