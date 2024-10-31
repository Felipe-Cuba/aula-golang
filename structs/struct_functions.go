package structs

import "math"

func (sqr Square) GetArea() float32 {
	return sqr.SideA * sqr.SideB
}

func (sqr Square) GetPerimeter() float32 {
	return (sqr.SideA * 2) + (sqr.SideB * 2)
}

func (sqr *Square) SetSides(sideA, sideB float32) {
	sqr.SideA = sideA
	sqr.SideB = sideB
}

func (circle Circle) GetArea() float32 {
	total := math.Pi * (math.Pow(float64(circle.Radius), 2))

	return float32(total)
}

func (circle Circle) GetPerimeter() float32 {
	total := 2 * math.Pi * circle.Radius

	return float32(total)
}

func (circle *Circle) SetRadius(value float32) {
	circle.Radius = value
}

func GetSquareByKey(key string) Square {
	return SquareMap[key]
}
