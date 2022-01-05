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

	fmt.Println("MENU")
	fmt.Println("____")
	fmt.Println("1 - Cappuchino")
	fmt.Println("2 - Latte")
	fmt.Println("4 - Americano")
	fmt.Println("5 - Mocha")
	fmt.Println("6 - Macchiato")
	fmt.Println("Q - Quit the program")

	fmt.Println(("Press any key in the keyboard. Press ESC to quit"))
	for {

		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		i, _ := strconv.Atoi(string(char))
		t := fmt.Sprintf("You chose %d", i)
		fmt.Println("You chose", char)
		fmt.Println(t)
		if char == 'q' || char == 'Q' {
			break
		}

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
