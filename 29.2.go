package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	i := 1
	for {
		select {
		case <-c:
			fmt.Println()
			fmt.Println("выхожу из программы")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Квадрат", i, "равен", i*i)
			i++
		}
	}
}
