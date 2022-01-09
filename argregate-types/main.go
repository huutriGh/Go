package main

import "fmt"

// aggregate types (array, struct)

// reference type (pointers, slices, maps, functions, channels)

// interface type

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func main() {

	// var myCar Car
	// myCar.NumberOfTires = 4
	// myCar.Luxury= false
	// myCar.Make = "Volkswagen"

	myCar := Car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "Volvo",
		Model:         "XC90",
		Year:          2019,
	}
	fmt.Printf("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)

	// var myString [3]string
	// myString[0] = "cat"
	// myString[1] = "dog"
	// myString[2] = "fish"
	// fmt.Println("First element in array is", myString[0])

}
