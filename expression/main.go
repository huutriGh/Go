package main

import "fmt"

func main() {
	age := 10
	name := "jack"
	rightHanded := true

	fmt.Printf("%s is %d years old. Right haned: %t", name, age, rightHanded)
	fmt.Println()

	ageInTenYears := age + 10

	fmt.Printf("In ten years, %s will be %d years old", name, ageInTenYears)
	fmt.Println()
	isTeenager := age >= 13

	fmt.Println(name, "is a teenager:", isTeenager)
}
