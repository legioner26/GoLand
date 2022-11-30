package student

import (
	"fmt"
)

type Student struct {
	name  string
	age   int
	grade int
}

func Put(name string, age int, grade int) *Student {
	return &Student{name, age, grade}
}

func Print(s Student) {
	fmt.Println(s.name, s.age, s.grade)
}

func GetName(s *Student) string {
	return s.name
}
