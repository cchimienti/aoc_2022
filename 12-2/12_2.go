/*
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
A/X = Rock 			1
B/Y = Paper			2
C/Z = Scissors		3

Win					6
Draw				3
Lost				0
______________________

Paper > Rock
Scissors > Paper
Rock > Scissors

- Check for win by comparing columns (calc draw first - nah), note column 2 value (in object struct),
tally winorloss and store ++

Your total score is the sum of your scores for each round.
The score for a single round is the score for the shape you selected
(1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round
(0 if you lost, 3 if the round was a draw, and 6 if you won).

//*

// Scoreboard
var round [2]string
var shape_tally int = 0
var total_tally int = 0
var fight_tally int = 0
var final_tally int = 0
var winner bool

func player_detect(shape string, Rock, Paper, Scissors arsenal.Player) arsenal.Player {
	if (shape == "A") || (shape == "X") {
		return Rock
	} else if (shape == "B") || (shape == "Y") {
		return Paper
	} else if (shape == "C") || (shape == "Z") {
		return Scissors
	}
	panic("what are you....")
}

func fight(opp arsenal.Player, me arsenal.Player) (bool, int) {
	shape_value := me.GetValue()
	shape_tally += shape_value

	// Draw
	if opp == me {
		fight_tally += 3
		//break?
		winner = false
		fmt.Printf("Draw! \n")
	}
	// W or L
	if opp.GetType() == "rock" {
		winner = me.AgainstRock()
	} else if opp.GetType() == "paper" {
		winner = me.AgainstPaper()
	} else if opp.GetType() == "scissors" {
		winner = me.AgainstScissors()
	}

	// Round Results
	if winner == true {
		fight_tally += 6
	}
	total_tally = shape_tally + fight_tally
	fmt.Printf("Value: %d + Winnings: %d = %d \n", shape_tally, fight_tally, total_tally)
	return winner, total_tally
}

func main1() {

	// Introduce Weapons
	Rock := arsenal.Player{Weapon: "rock", Value: 1, Fight_s: true, Fight_p: false}
	Paper := arsenal.Player{Weapon: "paper", Value: 2, Fight_s: false, Fight_r: true}
	Scissors := arsenal.Player{Weapon: "scissors", Value: 3, Fight_r: false, Fight_p: true}
	//fmt.Println(Rock, Paper, Scissors)

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
		opp, me, _ := strings.Cut(round, " ")
		fmt.Printf("%s -- %s \n", opp, me)

		// Detect Players
		fighter := player_detect(opp, Rock, Paper, Scissors)
		mii := player_detect(me, Rock, Paper, Scissors)

		// FIGHT
		win, total := fight(fighter, mii)
		fmt.Println(total)
		fmt.Println(win)

	}

	defer f.Close()
}

*/
