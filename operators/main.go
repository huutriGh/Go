package main

import (
	"fmt"
	"math"
)

func main() {
	//Multiplication
	// area =mr2

	var radius = 12.0
	area := math.Pi * radius * radius
	fmt.Println("area is", area)

	// iteger division

	half := 1 / 2
	fmt.Println("half with integer division:", half)
	halfFloat := 1.0 / 2.0
	fmt.Println("half float:", halfFloat)

	// squaring (raising to the power)

	badThreeSquared := 3 ^ 2
	fmt.Println("badThreeSquared:", badThreeSquared)
	goodThreeSquared := int(math.Pow(3.0, 2.0))
	fmt.Println("goodThreeSquare:", goodThreeSquared)

	//modulus

	remainder := 50 % 3

	fmt.Println("remaider", remainder)

}
