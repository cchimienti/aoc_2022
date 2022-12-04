package main

import (
	arsenal "12_2/pkg"
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Rock, Paper, Scissors
/*
A = Rock 			1
B = Paper			2
C = Scissors		3

X = Lost				0
Y = Draw				3
Z = Win					6
______________________

Paper > Rock
Scissors > Paper
Rock > Scissors

- check for Opp player and result, then compute needed player type
X + Y = Z
X + Z = Y

Your total score is the sum of your scores for each round.
The score for a single round is the score for the shape you selected
(1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round
(0 if you lost, 3 if the round was a draw, and 6 if you won).

*/

// Scoreboard
var round [2]string
var shape_tally int = 0
var total_tally int = 0
var fight_tally int = 0
var winner bool
var prediction string
var shape_value int

func player_detect(shape string, Rock, Paper, Scissors arsenal.Player) arsenal.Player {
	// Return everything about it in an array
	if shape == "A" {
		return Rock
	} else if shape == "B" {
		return Paper
	} else if shape == "C" {
		return Scissors
	}
	panic("what are you....")

}

func predict_fight(opp arsenal.Player, result string, RPS [3]arsenal.Player) { //(bool, int) {
	if result == "Y" { // Draw
		fight_tally += 3
		fmt.Printf("Draw! \n")
		prediction = opp.GetType()

	} else if result == "X" { // Lose
		prediction = opp.GetWin()
		fmt.Printf("Loser: %s \n", prediction)
		winner = false

	} else if result == "Z" { // Win
		prediction = opp.GetDefeat()
		fmt.Printf("Winner: %s \n", prediction)
		fight_tally += 6
		winner = true
	}

	// Values
	for _, field := range RPS {
		if prediction == field.GetType() {
			shape_value = field.GetValue()
			fmt.Printf("Value of my play - %d \n", shape_value)
		}
	}

	shape_tally += shape_value

	// Round Results
	total_tally = shape_tally + fight_tally
	fmt.Printf("Value: %d + Winnings: %d = %d \n", shape_tally, fight_tally, total_tally)
	//return winner, total_tally
}

func main() {

	// Introduce Weapons
	Rock := arsenal.Player{Weapon: "rock", Value: 1, Fight_s: true, Fight_p: false, Weak: "paper", Strong: "scissors"}
	Paper := arsenal.Player{Weapon: "paper", Value: 2, Fight_s: false, Fight_r: true, Weak: "scissors", Strong: "rock"}
	Scissors := arsenal.Player{Weapon: "scissors", Value: 3, Fight_r: false, Fight_p: true, Weak: "rock", Strong: "paper"}
	RPS := [3]arsenal.Player{Rock, Paper, Scissors}

	// Open File
	f, err := os.Open("Real_guide.txt")
	if err != nil {
		panic(err)
	}
	// bufio method (still need f)
	bf := bufio.NewScanner(f)
	//bf.Split(bufio.ScanWords)

	for bf.Scan() {
		round := bf.Text()
		opp, result, _ := strings.Cut(round, " ")
		fmt.Printf("%s -- %s \n", opp, result)

		// Detect Players
		fighter := player_detect(opp, Rock, Paper, Scissors)
		fmt.Printf("opp info: %d ", fighter.GetType())

		// FIGHT
		predict_fight(fighter, result, RPS)
		//fmt.Println(total)
		//fmt.Println(win)
		//fmt.Printf("my hand: %d \n", my_play)

	}

	defer f.Close()
}
