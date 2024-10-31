package structs

import "golang-aula/interfaces"

var SquareMap = map[string]Square{
	"rental": {
		SideA: 10,
		SideB: 20,
	},
	"square": {
		SideA: 15,
		SideB: 15,
	},
}

var square_rental = Square{
	SideA: 1.5,
	SideB: 2,
}

var square = Circle{
	Radius: 3,
}

var SquareArray = []interfaces.Figure{
	square_rental,
	square,
}
