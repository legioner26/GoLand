package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("12-1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileStat, _ := file.Stat()
	buf := make([]byte, fileStat.Size())
	_, err = file.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}
