package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("12-3.txt")
	if err := os.Chmod("12-3.txt", 0444); err != nil {
		fmt.Println(err)
	}
	file.WriteString("Текст")
	if err != nil {
		fmt.Println(err)
	}
	file.Close()

	file, err = os.Open("12-3.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
}
