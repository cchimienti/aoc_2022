package main

/*

Part II:
X controls the horizontal position of a sprite (3 pixels wide ###) - middle location


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

var SpritePositions [40]string
var Pixels []string
var Image [][]string

func AsyncNoop() {
	//fmt.Println(cycles)
	if cycles == 240 {
		Image = append(Image, Pixels)
	}
	s1, s2, s3 := DrawSprite(X)
	SpriteIntersect(s1, s2, s3)
	cycles++

}

func AsyncAddx(val int) (int, int) {
	for i := 0; i < 2; i++ {
		fmt.Printf("Cycles: %d \n", cycles)
		s1, s2, s3 := DrawSprite(X)
		SpriteIntersect(s1, s2, s3)
		cycles++
	}
	X += val
	return cycles, X
}

func DrawSprite(X int) (int, int, int) {
	if X <= 0 {
		return X - 1, X, X + 1
	} else if X > 38 {
		return X - 1, X, X + 1
	}
	SpritePositions[X-1] = "#"
	SpritePositions[X] = "#"
	SpritePositions[X+1] = "#"
	fmt.Println(SpritePositions)
	SpritePositions = [40]string{} //clear
	return X - 1, X, X + 1

}

func SpriteIntersect(sp1, sp2, sp3 int) {
	if cycles == 40 || cycles == 80 || cycles == 120 || cycles == 160 || cycles == 200 || cycles == 239 {
		Image = append(Image, Pixels)
		//fmt.Println(Pixels)
		Pixels = nil
	}
	if (cycles%40) == sp1 || (cycles%40) == sp2 || (cycles%40) == sp3 {
		Pixels = append(Pixels, "#")
		//fmt.Println(Pixels)
		//return true
	} else {
		Pixels = append(Pixels, ".")
		//fmt.Println(Pixels)
		//return false
	}
	//return false
}

func main() {

	cycles = 0
	X = 1

	// Open File
	f, err := os.Open("Real_program.txt")
	if err != nil {
		panic(err)
	}
	bf := bufio.NewScanner(f)

	for bf.Scan() { // already a for loop, may need to run longer tho
		// Read in input
		x := string(bf.Text())
		instr, add_str, _ := strings.Cut(x, " ")
		add_num, _ := strconv.Atoi(add_str)
		fmt.Println(instr, add_num)
		fmt.Printf("Start of Cycles: %d, X: %d \n", cycles, X)
		if instr == "noop" {
			AsyncNoop()

		} else if instr == "addx" {
			cycles, final = AsyncAddx(add_num)
		}
		//(val, numcheck)
		fmt.Printf("End of Cycle: %d, X: %d \n", cycles, X)
		//DrawSprite(X)
		//fmt.Println(Executions)
		//cycles++

	}
	for _, j := range Image {
		fmt.Printf("%s \n", j)
	}
	// Answer ~~ RJERPEFC
	defer f.Close()

}
