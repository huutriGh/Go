package main

import "fmt"

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal) Says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
	fmt.Println()
}

func (a *Animal) HowManyLegs() {

	fmt.Printf("A %s has %d legs", a.Name, a.NumberOfLegs)
	fmt.Println()
}

func main() {
	z := addTwoNumber(2, 4)
	fmt.Println(z)

	myTotal := sumMany(2, 3, 4, 5)
	fmt.Println(myTotal)

	var dog Animal
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.Says()

	cat := Animal{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
	}
	cat.Says()
	cat.HowManyLegs()
}

func addTwoNumber(x, y int) int {
	return x + y
}

// func addTwoNumber(x, y int) (sum int) {

// 	sum = x + y
// 	return
// }

func sumMany(num ...int) int {
	total := 0
	for _, x := range num {

		total = total + x

	}
	return total
}
