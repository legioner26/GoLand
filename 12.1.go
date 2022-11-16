package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	file, err := os.Create("12-1.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	counter := 1
	for {
		fmt.Println("Введите строку")
		var word string
		fmt.Scan(&word)
		if word == "exit" {
			fmt.Println("Завершаем программу")
			break
		}
		var b bytes.Buffer
		t := time.Now()
		timeWrite := t.Format("2006-01-02 15:04:05: ")
		b.WriteString(strconv.Itoa(counter) + ": " + timeWrite + word + "\n")
		_, err = file.Write(b.Bytes())
		if err != nil {
			fmt.Println(err)
			return
		}
		counter++
	}

}
