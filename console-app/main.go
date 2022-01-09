package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {

		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappuchino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("____")
	fmt.Println("1 - Cappuchino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program")

	fmt.Println(("Press any key in the keyboard. Press ESC to quit"))
	for {

		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if char == 'q' || char == 'Q' {
			break
		}
		i, _ := strconv.Atoi(string(char))
		t := fmt.Sprintf("You chose %s", coffees[i])

		fmt.Println(t)

	}

	fmt.Println("Program exiting...")
	// reader := bufio.NewReader(os.Stdin)
	// for {
	// 	fmt.Print("-> ")
	// 	userInput, _ := reader.ReadString('\n')

	// 	userInput = strings.Replace(userInput, "\r\n", "", -1)
	// 	userInput = strings.Replace(userInput, "\n", "", -1)

	// 	if userInput == "quit" {
	// 		break
	// 	} else {
	// 		fmt.Println(userInput)
	// 	}
	// }

}
