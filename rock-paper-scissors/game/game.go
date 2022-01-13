package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

func (g *Game) Rounds() {
	// use select to process in channles
	// Print to screen
	// Keep track of round number
	for {
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChan <- 1

		case msg := <-g.DisplayChan:
			fmt.Println(msg)
			g.DisplayChan <- ""
		}
	}

}

// clearScreen clears the screen
func (g *Game) ClearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
func (g *Game) PrintIntro() {
	g.DisplayChan <- `
Rock, paper & scissors
------------------------
Game is played for three rounds, and best of wins the game. Good luck!
`
	<-g.DisplayChan
}

func (g *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())
	playerValue := -1
	g.DisplayChan <- fmt.Sprintf(`
Round %d
--------
`, g.Round.RoundNumber)
	<-g.DisplayChan
	fmt.Print("Please enter rock, paper, scissors ->")
	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\r\n", "", -1)

	computerValue := rand.Intn(3)
	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}
	g.DisplayChan <- ""
	<-g.DisplayChan
	g.DisplayChan <- fmt.Sprintf("Player chose %s", strings.ToUpper(playerChoice))
	<-g.DisplayChan
	switch computerValue {
	case ROCK:
		g.DisplayChan <- "Computer chose ROCK"
		<-g.DisplayChan

	case PAPER:
		g.DisplayChan <- "Computer chose PAPER"
		<-g.DisplayChan
	case SCISSORS:
		g.DisplayChan <- "Computer chose SCISSORS"
		<-g.DisplayChan
	default:

	}
	fmt.Println()
	if playerValue == computerValue {
		g.DisplayChan <- "It's  a draw"
		<-g.DisplayChan
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}
		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}
		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}
		default:

			g.DisplayChan <- "Invalid choise"
			<-g.DisplayChan
			return false
		}
	}
	return true
}

func (g *Game) computerWins() {

	g.Round.ComputerScore++
	g.DisplayChan <- "Computer wins!"
	<-g.DisplayChan
}
func (g *Game) playerWins() {

	g.Round.PlayerScore++
	g.DisplayChan <- "Player wins!"
	<-g.DisplayChan

}

func (g *Game) PrintSummary() {

	fmt.Println("Final score")
	fmt.Println("------------")
	fmt.Printf("Player: %d/3, computer %d/3", g.Round.PlayerScore, g.Round.ComputerScore)
	fmt.Println()
	if g.Round.PlayerScore > g.Round.ComputerScore {
		fmt.Println("Player wins game!")

	} else {
		fmt.Println("computer wins game!")
	}
}
