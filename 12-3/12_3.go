package main

/*
	Find the occuring item type in two strings of the same length (can be > 1)
	Then rank by priority and sum that total

	Lowercase item types a through z have priorities 1 through 26.
	Uppercase item types A through Z have priorities 27 through 52.


	a - z 				1 - 26
	A - Z	 			27- 52
	________________________________________


	Looping through the alphabet in Go seems to be easiest / most efficient to use a rune (aka int32)
*/

import (
	"bufio"
	"fmt"
	"os"
)

var matched []string
var tally int
var group []string
var bigger_boot [][]string

// simple func - again... Get intersection of two now three strings
func GetIntersect(r1, r2, r3 string) (inter rune) {
	q2, q3 := false, false
	hashed := make(map[string]bool)
	hashed_2 := make(map[string]bool)
	for _, e := range r1 {
		hashed[string(e)] = true
	}
	for _, e := range r2 {
		if hashed[string(e)] {
			hashed_2[string(e)] = true
			q2 = true
		}
	}
	for _, e := range r3 {
		if hashed[string(e)] && hashed_2[string(e)] {
			q3 = true
			if q2 == true && q3 == true {
				fmt.Printf("matched -- %s \n", string(e))
				return e
			}
		}
	}
	return
}

func main() {
	// Open File
	f, err := os.Open("Real_rucksack.txt")
	if err != nil {
		panic(err)
	}
	bf := bufio.NewScanner(f)

	// New Alphabet Priority Map
	pm := make(map[string]rune)
	for r := 'a'; r <= 'z'; r++ {
		pm[string(rune(r))] = (r - 96)
	}
	for s := 'A'; s <= 'Z'; s++ {
		pm[string(rune(s))] = (s - 38)
	}
	fmt.Println(pm)

	for bf.Scan() {
		x := bf.Text()
		//fmt.Println(line_count)
		//fmt.Printf("incoming: %s \n", x)
		//fmt.Printf("lgroup: %d \n", len(group))

		if len(group) < 3 {
			group = append(group, x) // last line is true here, but conditional breaks
		}
		if len(group) == 3 {
			bigger_boot = append(bigger_boot, group)
			group = nil
			//group = append(group, x)
		}

		//compart_size := len(x) / 2
		//first, second := x[:compart_size], x[compart_size:]
		//intersect := GetIntersect(string(first), string(second))
		//for _, n := range bigger_boot {
		//	fmt.Printf("{%s} \n", n)
		//}
		//intersect := GetIntersect()
		//	matched = append(matched, string(intersect))
		//fmt.Println(group)

	}

	for _, elf_group := range bigger_boot {
		//fmt.Printf("%s \n", elf_group)
		r1 := elf_group[0]
		r2 := elf_group[1]
		r3 := elf_group[2]
		fmt.Printf("%s, %s, %s ~~ \n ", r1, r2, r3)
		intersect := GetIntersect(r1, r2, r3)
		//fmt.Println(intersect)
		matched = append(matched, string(intersect))
		//fmt.Println(matched)

	}

	defer f.Close()

	// Match Intersection to Priority Map & Tally
	fmt.Println(matched)
	//fmt.Println((bigger_boot))
	for _, v := range matched {
		tally += int(pm[v])
		fmt.Println(tally)
	}
}
