package main

import (
	"fmt"
)

func main() {
	//Multiplication
	// area =mr2

	// var radius = 12.0
	// area := math.Pi * radius * radius
	// fmt.Println("area is", area)

	// // iteger division

	// half := 1 / 2
	// fmt.Println("half with integer division:", half)
	// halfFloat := 1.0 / 2.0
	// fmt.Println("half float:", halfFloat)

	// // squaring (raising to the power)

	// badThreeSquared := 3 ^ 2
	// fmt.Println("badThreeSquared:", badThreeSquared)
	// goodThreeSquared := int(math.Pow(3.0, 2.0))
	// fmt.Println("goodThreeSquare:", goodThreeSquared)

	// //modulus

	// remainder := 50 % 3

	// fmt.Println("remaider", remainder)

	// precedence
	// multiplication
	// a := 12.0 * 3.0 / 4.0
	// b := (12.0 * 3.0) / 4.0
	// c := 12 * (3.0 / 4.0)

	// fmt.Println("a", a, "b", b, "c", c)

	// // integer devision
	// unclear := 12 * (3 / 4)
	// fmt.Println("unclear", unclear)

	// // parenthesis
	// f := 12.0 / 3.0 / 4.0

	// fmt.Println("f", f)

	// f = 12.0 / (3.0 / 4.0)
	// fmt.Println("f", f)

	//does one number divide exactly into another?
	// x := 12
	// y := 3

	// if x%y == 0 {
	// 	fmt.Println(y, "divides exactly into", x)
	// } else {
	// 	fmt.Println(y, "does not divid exactly into", x)
	// }

	for m := 0; m <= 12; m++ {

		fmt.Println("The month after", m, "is", m%12+1)

	}
}
