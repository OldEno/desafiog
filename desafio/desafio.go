package main

import (
	"fmt"
	"time"
)

func pingar(c1, c2 chan string) {
	for i := 0; ; i++ {
		c1 <- "ping"
		time.Sleep(time.Second)
		c2 <- "pong"
		time.Sleep(time.Second)
	}
}

func imprimir(c1, c2 chan string) {
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go pingar(c1, c2)
	go imprimir(c1, c2)

	var entrada string
	fmt.Scanln(&entrada)
}