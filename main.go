package main

import (
	"fmt"
	"golang-aula/functions"
	"time"
)

func main() {
	defer fmt.Println("terminou")

	chanel := make(chan int)

	go setIntenger(chanel, 100)

	go printValues(chanel)

	time.Sleep(10 * time.Second)
}

func setIntenger(ch chan int, value int) {
	ch <- value
}

func printValues(ch chan int) {
	intenger := <-ch

	go functions.PrintEven(intenger)
	go functions.PrintOdd(intenger)

	time.Sleep(10 * time.Second)
}
