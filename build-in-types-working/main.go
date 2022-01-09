package main

import "log"

// basic type (numbers, string, boolean)

var myInt int

var myUnit uint

var myFloat float32

var myFloat64 float64

// aggregate types (array, struct)

// reference type (pointers, slices, maps, functions, channels)

// interface type

func main() {
	myInt = 10
	myUnit = 20
	myFloat = 10.1
	myFloat64 = 100.1

	log.Println(myInt, myUnit, myFloat, myFloat64)

	myString := "Tri"
	log.Println(myString)
	myString = "Trevor"
	var myBool = true
	log.Println(myBool)
}
