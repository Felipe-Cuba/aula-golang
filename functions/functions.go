package functions

import (
	"fmt"
	"golang-aula/interfaces"
	"strconv"
	"sync"
	"time"
)

func Soma(n1, n2 int) int {
	return n1 + n2
}

func GetVariables() (int, int) {
	return privada, Publica
}

func SumArea(figArray []interfaces.Figure) float32 {

	var totalArea float32

	for _, fig := range figArray {
		totalArea += fig.GetArea()
	}

	return totalArea
}

func SumPerimeters(figArray []interfaces.Figure) float32 {
	var totalPerimeter float32

	for _, fig := range figArray {
		totalPerimeter += fig.GetPerimeter()
	}

	return totalPerimeter

}

func SetIntengerChanel(ch chan int, value int) {
	ch <- value
}

func PrintIntengerChanel(ch chan int) {
	intenger := <-ch

	fmt.Println(intenger)
}

func PrintValueWaitGroup(wg *sync.WaitGroup, value string) {
	defer wg.Done()

	time.Sleep(time.Second)

	intTen, err := strconv.Atoi(value)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(intTen)
}

func PrintEven(total int) {
	for i := 0; i < total; i += 2 {
		fmt.Println(i)
	}

}

func PrintOdd(total int) {

	for i := 1; i < total; i += 2 {
		fmt.Println(i)
	}
}
