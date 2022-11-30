package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name  string
	age   int
	grade int
}

func (s Student) info() string {
	return s.name + " " + strconv.Itoa(s.age) + " " + strconv.Itoa(s.grade)
}

func main() {
	studentMap := make(map[string]*Student, 0)

	fmt.Println("Пожалуйста, введите имя, возраст и курс студентов с переносом строки:")
	fmt.Println("---------------------")

	var counter = 1
	var in = bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		line, err := in.ReadString('\n')

		if err == io.EOF {
			fmt.Print("-\n")
			fmt.Println("---------------------")
			fmt.Println("Конец ввода! Список студентов:")
			break
		}

		lineFields := strings.Fields(line)

		if len(lineFields) < 3 {
			fmt.Print("Необходимо ввести имя, возраст и грейд! Пожалуйста, попробуйте снова...\n")
			continue
		}

		studentName := lineFields[0]
		studentAge, errAge := strconv.Atoi(lineFields[1])
		studentGrade, errGrade := strconv.Atoi(lineFields[2])

		if errAge != nil || errGrade != nil {
			fmt.Print("Ошибка при обработке возраста студента и его грейда! Пожалуйста, попробуйте снова...\n")
			continue
		}

		student := Student{
			name:  studentName,
			age:   studentAge,
			grade: studentGrade,
		}

		if _, err := get(studentMap, student.name); err != nil {
			newStudent(studentMap, &student)
		} else {
			fmt.Print("Студент с таким именем уже есть в хранилище! Попробуйте снова...\n")
		}
	}

	for _, value := range studentMap {
		fmt.Printf("%v. %s\n", counter, value.info())
		counter++
	}
}

func newStudent(list map[string]*Student, value *Student) (int, error) {
	list[value.name] = value

	if list[value.name] == nil {
		return -1, errors.New("Ошибка добавления данных в хранилище!")
	} else {
		return 0, nil
	}
}

func get(list map[string]*Student, name string) (*Student, error) {
	if list[name] == nil {
		return nil, errors.New("Студент не найден в хранилище!")
	} else {
		return list[name], nil
	}
}
