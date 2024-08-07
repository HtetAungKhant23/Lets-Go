package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome from rock, paper, scissor game")
	rounds := 3
	winCounter := 0

	for range rounds {
		var playerChoice, computerChoice string
		fmt.Print("Enter your choice (Rock, Paper, Scissor): ")
		fmt.Scanln(&playerChoice)

		switch random := rand.Intn(rounds); random {
		case 0:
			computerChoice = "Rock"
		case 1:
			computerChoice = "Paper"
		case 2:
			computerChoice = "Scissor"
		}

		fmt.Println("Computer choose", computerChoice)

		switch {
		case playerChoice == computerChoice:
			fmt.Println("Tie!")
		case playerChoice == "Rock" && computerChoice == "Scissor",
			playerChoice == "Paper" && computerChoice == "Rock",
			playerChoice == "Scissor" && computerChoice == "Paper":
			fmt.Println("Player win this round")
			winCounter++
		default:
			fmt.Println("Computer win this round")
			winCounter--
		}
	}

	if winCounter > 0 {
		fmt.Println("Congratulation. You win", winCounter, "times")
	} else if winCounter == 0 {
		fmt.Println("Congratulation. The game is Tie")
	} else {
		fmt.Println("Game over")
	}
}
