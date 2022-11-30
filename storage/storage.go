package storage

import (
	"fmt"
	"skillbox/student"
)

var mapStudent map[string]*student.Student = make(map[string]*student.Student, 0)

func AddStud(st *student.Student) {
	mapStudent[student.GetName(st)] = st
}

func PrintStudent() {
	fmt.Println()
	fmt.Println("Студенты из хранилища:")
	for _, v := range mapStudent {
		student.Print(*v)
	}
}
