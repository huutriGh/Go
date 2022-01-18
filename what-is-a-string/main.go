package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println()
	// name := "Hello world"
	// fmt.Println("String:", name)
	// fmt.Println()
	// fmt.Println("Bytes")

	// for i := 0; i < len(name); i++ {
	// 	fmt.Printf("%x ", name[i])
	// }

	// fmt.Println()
	// fmt.Println()
	// fmt.Println("Index\tRune\tString")
	// for x, y := range name {

	// 	fmt.Println(x, "\t", y, "\t", string(y))
	// }
	// fmt.Println()

	// name = "Γειά σου Κόσμε"
	// for x, y := range name {

	// 	fmt.Println(x, "\t", y, "\t", string(y))
	// }
	// fmt.Println()
	// fmt.Println()
	// fmt.Println("Three ways to concatnate strings")

	// h := "Hello, "
	// w := "world"

	// // using + not very efficient
	// myString := h + w
	// fmt.Println(myString)
	// fmt.Println()

	// // using fmt - more efficient
	// myString = fmt.Sprintf("%s%s", h, w)
	// fmt.Println(myString)
	// fmt.Println()

	// // using stringbuilder - very efficient

	// var sb strings.Builder
	// sb.WriteString(h)
	// sb.WriteString(w)
	// fmt.Println(sb.String())

	// fmt.Println()

	// name = "ABCDEFGHIKLMNOPQRSTUYWXZY"
	// fmt.Println("Getting a substring")
	// fmt.Println(name[10:13])

	// courseName := "Learn Go for Beginers Crash Course"

	// fmt.Println(string(courseName[0]))
	// fmt.Println(string(courseName[6]))

	// for i := 0; i <= 21; i++ {

	// 	fmt.Print(string(courseName[i]))

	// }
	// fmt.Println()
	// for i := 13; i <= 21; i++ {

	// 	fmt.Print(string(courseName[i]))

	// }
	// fmt.Println()

	// fmt.Println("Length of courseName is", len(courseName))

	// var mySlice []string
	// mySlice = append(mySlice, "one")
	// mySlice = append(mySlice, "two")
	// mySlice = append(mySlice, "three")

	// fmt.Println("mySlice has", len(mySlice), "elements")
	// fmt.Println("the last element in mySlice is", mySlice[len(mySlice)-1])

	// courses := []string{
	// 	"Learn Go for Beginers Crash Course",
	// 	"Learn Java for Beginers Crash Course",
	// 	"Learn Python for Beginers Crash Course",
	// 	"Learn C for Beginers Crash Course",
	// }

	// for _, x := range courses {
	// 	if strings.Contains(x, "Go") {
	// 		fmt.Println("Go is found in", x, "and index is", strings.Index(x, "Go"))
	// 	}
	// }
	// newString := "Go is a great programming language. Go for it!"
	// fmt.Println(strings.HasPrefix(newString, "Go"))
	// fmt.Println(strings.HasPrefix(newString, "Python"))
	// fmt.Println(strings.HasSuffix(newString, "!"))
	// fmt.Println(strings.Count(newString, "Go"))
	// fmt.Println(strings.Index(newString, "Python"))
	// fmt.Println(strings.LastIndex(newString, "Go"))

	// newString := "Go is a great programming language. Go for it!"

	// newString = strings.Replace(newString, "Go", "Golang", -1)

	// fmt.Println(newString)

	// // string comparation

	// if "Alpha" > "Absolute" {
	// 	fmt.Println("A is greater than B")
	// } else {
	// 	fmt.Println("A is not greater than B")
	// }

	// badMail := " me@here.com "

	// badMail = strings.TrimSpace(badMail)

	// fmt.Printf("=%s=", badMail)
	// fmt.Println()

	// str := "alpha alpha alpha alpha alpha"
	// str = replaceNth(str, "alpha", "beta", 3)
	// fmt.Println(str)

	myString := "This is a clear EXAMPLE of why we search in one case only"

	searchString := strings.ToLower(myString)
	if strings.Contains(searchString, "this") {
		fmt.Println("Found it!")
	} else {
		fmt.Println("Did not find it!")
	}

	// other case function

	fmt.Println(strings.ToLower(myString))
	fmt.Println(strings.ToUpper(myString))
	fmt.Println((strings.Title(strings.ToLower(myString))))

}

// func replaceNth(s, old, new string, n int) string {
// 	i := 0
// 	for j := 1; j <= n; j++ {
// 		x := strings.Index(s[i:], old)
// 		if x < 0 {
// 			break
// 		}
// 		//  have found it

// 		i = i + x
// 		if j == n {
// 			return s[:i] + new + s[i+len(old):]
// 		}
// 		i += len(old)
// 	}
// 	return s
// }
