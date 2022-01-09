package main

import "fmt"

// reference type (pointers, slices, maps, functions, channels)

// interface type

func main() {

	x := 10
	myFirstPointer := &x
	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is", myFirstPointer)
	*myFirstPointer = 15
	fmt.Println("x is now", x)

	changeValueOfPointer(&x)
	fmt.Println("After function called, x is now", x)

}

func changeValueOfPointer(num *int) {
	*num = 25
}
