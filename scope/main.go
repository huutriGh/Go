package main

import "myapp/packageone"

var myvar = "I am myVar"

func main() {
	// variables
	// declare a pagekage level variable for the main package name myvar

	// declare a block variable for the main function call blockvar

	var blockVar = "I am blockvar"

	// declare a package level variable in the package one package name PackageVar

	// create an exported fuction in packageone one called PrintMe

	// in the main function, print out the values of myvar, blockVar, and Packagevar on one line using he PrintMe
	// function in packageone
	packageone.PrintMe(myvar, blockVar)

}
