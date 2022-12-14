package main

//package main

/*

addx V takes two cycles to complete. After two cycles, the X register is increased by the value V. (V can be negative.)
noop takes one cycle to complete. It has no other effect.

X = 1 to begin

Find the signal strength during the 20th, 60th, 100th, 140th, 180th, and 220th cycles

*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var X int
var cycles int
var final int
var signal_strength int

// var val int
// var numcheck int
var Executions []int

func AsyncNoop() {
	fmt.Println(cycles)
	//fmt.Printf("cycle Noop \n")
	cycles++

}

func AsyncAddx(val int) (int, int) { //(X int, executenow int)[]int { // Note which cycle to execute here
	for i := 0; i < 2; i++ {
		fmt.Println(cycles)
		if cycles == 20 {
			//signal_strength += X * 20
			fmt.Println(cycles, X)
		} else if cycles == 60 {
			//signal_strength += X * 60
			fmt.Println(cycles, X)
		} else if cycles == 220 {
			//signal_strength += X * 100
			fmt.Println(cycles, X)
			//fmt.Printf("Cycles: %d, X: %d \n", cycles, X)
			//return cycles, X
		}
		cycles++
	}
	X += val
	return cycles, X
	//cycles += 2
	//executenow = addxcycle + 2
	//return make([]int, val, executenow)
	//return val, executenow

}

func CheckCycleExec(Executions []int) { //Pass in correct execute cycle - should check every cycle, check list of marked spots
	//fmt.Println(cycles, executenow)
	//for k, val := range Executions {
	//fmt.Println(k, val)
	//if k%2 == 0 {
	//	fmt.Printf("Check urself %d -- %d \n", cycles, Executions[k+1])
	//	if cycles == Executions[k+1] && cycles != 0 {
	//		fmt.Printf("Cycle Reached --- X updated \n")
	//		X += val
	//		fmt.Println(X)

	//		} else {
	//			fmt.Printf("skippy skip \n")
	//		}

	//	}
	//}
}

func main() {
	// Open File
	f, err := os.Open("Real_program.txt")
	if err != nil {
		panic(err)
	}
	bf := bufio.NewScanner(f)

	cycles = 1
	X = 1
	//val = 0
	//numcheck = 0

	for bf.Scan() { // already a for loop, may need to run longer tho
		// Read in input
		x := string(bf.Text())
		instr, add_str, _ := strings.Cut(x, " ")
		add_num, _ := strconv.Atoi(add_str)
		//fmt.Printf("Start of Cycles: %d, X: %d \n", cycles, X)
		if cycles == 20 {
			signal_strength += X * 20
			fmt.Println(cycles, signal_strength)
		} else if cycles == 60 {
			signal_strength += X * 60
			fmt.Println(cycles, signal_strength)
		} else if cycles == 100 {
			signal_strength += X * 100
			fmt.Println(cycles, signal_strength)
		} else if cycles == 140 {
			signal_strength += X * 140
			fmt.Println(cycles, signal_strength)
		} else if cycles == 180 {
			signal_strength += X * 180
			fmt.Println(cycles, signal_strength)
		} else if cycles == 220 {
			signal_strength += X * 220
			fmt.Println(cycles, signal_strength)
		}
		//CheckCycleExec(Executions)
		if instr == "noop" {
			AsyncNoop()
			//CheckCycleExec(AsyncAddx(add_num, cycles))

		} else if instr == "addx" {
			//fmt.Printf("++ %d \n", add_num)
			cycles, final = AsyncAddx(add_num)
			//fmt.Println(cycles, final)
			//val, numcheck := AsyncAddx(add_num, cycles) //mark cycle where instr is called, execute when cycles is that
			//Executions = append(Executions, val, numcheck)
			// push to list
			//CheckCycleExec(val, numcheck)
		}
		//CheckCycleExec(val, numcheck)

		//fmt.Println(Executions)
		//cycles++

	}
	fmt.Println(signal_strength)
	defer f.Close()

}
