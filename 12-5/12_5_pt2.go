package main

/*
	Implementing Stacks basically
	- Parse first, then implement stacks based on input

	This one got away from me tbh - became more like an excercie in Go

	PART II =

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
var extra_carry []string

func build_stacks(stack_column int, queue []string) stack.Stack {
	var tower stack.Stack
	for _, j := range queue {
		tower.Push(j)
	}
	return tower
}
func follow_instructions(Crates []stack.Stack, instructions []string) {
	for i := 0; i < len(instructions); i += 3 {
		many, _ := strconv.Atoi(instructions[i])
		orgin_stack, _ := strconv.Atoi(instructions[i+1])
		dest_stack, _ := strconv.Atoi(instructions[i+2])
		fmt.Println(many, orgin_stack, dest_stack)
		extra_carry = nil
		for j := 0; j < many; j++ {
			if many == 1 {
				top, _ := Crates[orgin_stack-1].Pop()
				Crates[dest_stack-1].Push(top)
				fmt.Println(j, Crates)
			} else {
				top, _ := Crates[orgin_stack-1].Pop()
				extra_carry = append(extra_carry, top)
				fmt.Printf("carried over: %s \n", extra_carry)
				if len(extra_carry) == many {
					extra := build_stacks(many, extra_carry) // Create temp additional Stack to put carried over objects back in
					// ew - triple-nested for loop is definitely not needed....
					for _, v := range extra_carry {
						top, _ := extra.Pop()
						Crates[dest_stack-1].Push(top)
						fmt.Println(v)
					}
				}
				fmt.Println(j, Crates)
			}
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

			if strings.Contains(x, "move") { // condition for move instructions {
				instructions = append(instructions, column...)
			}
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
	defer f.Close()

}
