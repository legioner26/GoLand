package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"skillbox/storage"
	"skillbox/student"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите имя студента возраст и курс через пробел")
	fmt.Println("Для завершения введите комбинацию клавишь Ctrl+D")

	var in = bufio.NewReader(os.Stdin)

	for {
		str, err := in.ReadString('\n')

		if err == io.EOF {
			break
		}
		addNewStudent(str)
	}

	storage.PrintStudent()
}

func addNewStudent(str string) {
	name, age, grade := getContent(str)
	storage.AddStud(student.Put(name, age, grade))
}

func getContent(str string) (name string, age int, grade int) {
	parseStr := strings.Fields(str)

	name = parseStr[0]
	age, _ = strconv.Atoi(parseStr[1])
	grade, _ = strconv.Atoi(parseStr[2])

	return
}
