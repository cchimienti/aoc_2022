package main

//package main

/*
	Implementing Stacks basically
	- Parse first, then implement stacks based on input

	This one got away from me tbh - became more like an excercie in Go

*/

import (
	stack "12_5/pkg"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var ordered_keys []int
var Crate stack.Stack
var Crates []stack.Stack
var instructions []string

func build_stacks(stack_column int, queue []string) stack.Stack {
	var tower stack.Stack
	for _, j := range queue {
		tower.Push(j)
	}
	return tower
}
func follow_instructions(Crates []stack.Stack, instructions []string) {
	for i := 0; i < len(instructions); i += 3 {
		//fmt.Println(instructions[i])
		many, _ := strconv.Atoi(instructions[i])
		orgin_stack, _ := strconv.Atoi(instructions[i+1])
		dest_stack, _ := strconv.Atoi(instructions[i+2])
		fmt.Println(many, orgin_stack, dest_stack)
		for j := 0; j < many; j++ {
			top, _ := Crates[orgin_stack-1].Pop()
			Crates[dest_stack-1].Push(top)
			//fmt.Printf("pushing... \n")
			//pushstack.Push(top)
			fmt.Println(j, Crates)
		}
	}
}

func main() {
	ha_queue := make(map[int][]string)
	// Open File
	f, err := os.Open("Real_supply.txt")
	if err != nil {
		panic(err)
	}
	bf := bufio.NewScanner(f)

	for bf.Scan() {
		x := bf.Text()
		fmt.Println(x)
		numeric := regexp.MustCompile(`\d`).MatchString(x)

		// Parse Stacks
		if numeric == false {
			mark_next := false
			for k, v := range x { // multiples of 4 + 1
				//fmt.Printf("%d -- %s\n", k, string(v))
				if mark_next == true && string(v) != " " {
					ha_queue[k] = append([]string{string(v)}, ha_queue[k]...)
				}
				if k%4 == 0 {
					mark_next = true
				} else {
					mark_next = false
				}
			}

			// Parse nums
		} else {
			column := regexp.MustCompile("[0-9]+").FindAllString(x, -1)
			//fmt.Println(column)

			if strings.Contains(x, "move") { // condition for move instructions {
				//many, orgin_stack, dest_stack := column[0], column[1], column[2]
				instructions = append(instructions, column...)
				//fmt.Println(many, orgin_stack, dest_stack)
			} /*else {
				max_columns, _ = strconv.Atoi(column[len(column)-1]) // number of columns
				fmt.Printf("~%d~ \n", max_columns)

			} */

		}

	}
	fmt.Println(ha_queue)

	// maps in go aren't ordered -- hacky...
	ordered_keys := make([]int, 0, len(ha_queue))
	for k := range ha_queue {
		ordered_keys = append(ordered_keys, k)
	}
	sort.Ints(ordered_keys)
	// iterate by sorted keys
	for k, v := range ordered_keys {
		//fmt.Println(k+1, v, ha_queue[v])
		m := build_stacks(k+1, ha_queue[v]) // build a new stack, pass in map to store, max amount, and value
		fmt.Println(m)
		Crates = append(Crates, m) // Create ordered array of Stacks - 3 in this example
	}

	follow_instructions(Crates, instructions)
	//fmt.Println(Crates)
	defer f.Close()

}
