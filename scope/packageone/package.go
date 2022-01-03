package packageone

import "fmt"

var PackageVar = "I am PackageVar"

func PrintMe(myvar, blockVar string) {
	fmt.Println(myvar, blockVar, PackageVar)
}
