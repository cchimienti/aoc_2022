package main

/*
	Treetop Tree Houses
	- Create a grid then calculate inner 'visibility' of trees based on tree heights; like an elevation map visualized
	create a 2d array 'visible' to store bools per index

	Second Part is to calculate a Scenic Score of how many trees can be seen
		Change from outside in to inside looking out
*/

import (
	"bufio"
	"fmt"
	"os"
)

var TreeHouse [][]int
var treerow []int
var rows int

var NonEdges [][]bool
var nonrow []bool

var Visbility [][]bool
var visCount int

var HighestScore int
var SecondHighest int

var columns int

// Need to only check for falsies, return true by default

// N, E, S, W
func CheckSouth(x, y int) (bool, int, int, int) { // North then South
	//for xv := rows - 1; xv >= 0; xv-- { // From Below!
	for xv := 0; xv < rows; xv++ { //As Above
		if xv > x { // Down
			if TreeHouse[xv][y] >= TreeHouse[x][y] {
				fmt.Printf("Tall Tree Found South \n")
				dist := xv - x // Maybe store distance between current and found tree here as viewing distance
				fmt.Printf("dist South: %d \n", dist)
				return false, x, y, dist
			}
		}
	}
	// Current is the tallest
	//fmt.Printf("Coast is clear S \n")
	return true, x, y, columns - x - 1 // if no trees encountered - calculate dist to

}

func CheckNorth(x, y int) (bool, int, int, int) {
	for xv := rows - 1; xv >= 0; xv-- { // From Below
		//for xv := 0; xv < rows; xv++ { //From Above
		//fmt.Println(xv, x)
		if xv < x { // Up
			if TreeHouse[xv][y] >= TreeHouse[x][y] {
				fmt.Printf("Tall Tree Found North \n")
				dist := x - xv
				fmt.Printf("dist North: %d \n", dist)
				return false, x, y, dist
			}
		}

	}
	// Current is the tallest
	//fmt.Printf("Coast is clear N \n")
	return true, x, y, x
}

func CheckLeft(x, y int) (bool, int, int, int) {
	for yv := y - 1; yv >= 0; yv-- {
		fmt.Println(yv, y)
		if yv < y { // Left
			if TreeHouse[x][yv] >= TreeHouse[x][y] {
				fmt.Printf("Tall Tree Found Left \n")
				dist := y - yv
				fmt.Printf("dist Left: %d \n", dist)
				return false, x, y, dist
			}
		}

	}
	// Current is the tallest
	//fmt.Printf("Coast is clear L \n")
	return true, x, y, y
}

func CheckRight(x, y int) (bool, int, int, int) {
	for yv := 0; yv < rows; yv++ {
		if yv > y { // Right
			if TreeHouse[x][yv] >= TreeHouse[x][y] {
				fmt.Printf("Tall Tree Found Right \n")
				dist := yv - y
				fmt.Printf("dist Right: %d \n", dist)
				return false, x, y, dist
			}
		}

	}
	// Current is the tallest
	//fmt.Printf("Coast is clear R \n")
	return true, x, y, rows - 1 - y
}

func MarkVisible(check bool, x, y int) {
	if check == true {
		Visbility[x][y] = false
		fmt.Println(Visbility)

	}
}

func GetScenicScore(N, S, E, W int) (int, int) {
	score := N * S * E * W
	//fmt.Println(score)
	if score > HighestScore {
		SecondHighest = HighestScore
		HighestScore = score
	}
	return score, HighestScore
}

//func CalculateVisible()

func main() {
	// Open File
	f, err := os.Open("Real_tree_grid.txt")
	if err != nil {
		panic(err)
	}
	bf := bufio.NewScanner(f)

	for bf.Scan() {
		//Build a grid, line by line - TreeHouse
		tree_grid := string(bf.Text())
		rows = len(tree_grid)
		for x, tree := range tree_grid {
			tree_height := int(tree - '0')
			treerow = append(treerow, tree_height) // get array to push into TreeHouse - let it rain
			if x == 0 || x == rows-1 {             // Edge Columns
				nonrow = append(nonrow, false) // let it rain
			} else {
				nonrow = append(nonrow, true)
			}
		}
		TreeHouse = append(TreeHouse, treerow)
		columns++
		NonEdges = append(NonEdges, nonrow)
		treerow = nil // clear it out
		nonrow = nil
	}

	// Define Edge Rows by visbility - first and last rows, & first and last indexes in fed line
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if i == 0 || i == rows-1 {
				NonEdges[i][j] = false
			}

		}

	}

	defer f.Close()

	// Calculations
	Visbility = NonEdges
	//fmt.Println(Visbility)

	// Iterate on indexes where NonEdges = true
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if NonEdges[i][j] == true {
				fmt.Println(i, j)
				//MarkVisible(CheckNorth(i, j))
				//MarkVisible(CheckSouth(i, j))
				//MarkVisible(CheckLeft(i, j))
				//MarkVisible(CheckRight(i, j))
				_, _, _, N := CheckNorth(i, j)
				_, _, _, E := CheckRight(i, j)
				_, _, _, S := CheckSouth(i, j)
				_, _, _, W := CheckLeft(i, j)
				//fmt.Println(W)
				fmt.Println(N, E, S, W)
				_, final := GetScenicScore(N, S, E, W)
				fmt.Println(final)
				fmt.Printf("\n")
			}
		}
	}

	// lazily calculate all the falses (aka trues)
	/*
		for i2 := 0; i2 < rows; i2++ {
			for j2 := 0; j2 < rows; j2++ {
				if Visbility[i2][j2] == false {
					visCount++
				}
			}
		}
	*/
	//fmt.Println(Visbility)
	//fmt.Println(visCount)

}
