package main

import (
	"fmt"
	"golang-aula/structs"
)

func main() {
	array_sqrs := []structs.Square{
		{
			SideA: 1.5,
			SideB: 3,
		},
		{
			SideA: 3,
			SideB: 3,
		},
	}

	fmt.Println(SumArea(array_sqrs[0], array_sqrs[1]))
}

func SumArea(sqr1, sqr2 structs.Square) float32 {
	return sqr1.GetArea() + sqr2.GetArea()
}
