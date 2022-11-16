package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	var b bytes.Buffer
	counter := 1
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите строку")
		text, _ := reader.ReadString('\n')
		value := strings.Fields(text)

		if value[0] == "exit" {
			fmt.Println("Завершаем запись")
			break
		}
		t := time.Now()
		timeStamp := t.Format("2006-01-02 15:04:05: ")

		var message []byte
		message = []byte(strconv.Itoa(counter) + ": " + timeStamp + text)

		b.Write(message)
		counter++
	}
	// запись из буффера в файл/создание файла, допуск
	err := os.WriteFile("12-4.txt", b.Bytes(), 0777)
	if err != nil {
		panic(err)
	}

	file, err := os.Open("12-4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// запись данных из файлв в buff
	buf, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Выводим содержимое файла")
	fmt.Println(string(buf))
}
