package main

/*
	for example, suppose you receive the following datastream buffer:

	mjqjpqmgbljsphdztnvjfqwrcgsmlb
	After the first three characters (mjq) have been received, there haven't been enough characters received yet to find the marker.
	The first time a marker could occur is after the fourth character is received, making the most recent four characters mjqj.
	Because j is repeated, this isn't a marker.

	The first time a marker appears is after the seventh character arrives. Once it does, the last four characters received are jpqm, which are all different.
	In this case, your subroutine should report the value 7, because the first start-of-packet marker is complete after 7 characters have been processed.

	Here are a few more examples:

	bvwbjplbgvbhsrlpgdmjqwftvncz: first marker after character 5
	nppdvjthqldpwncqszvftbrmjlhg: first marker after character 6
	nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg: first marker after character 10
	zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw: first marker after character 11

*/

import (
	"bufio"
	"fmt"
	"os"
)

func checkfordup14(scoped string) bool {
	hash := make(map[rune]bool)
	for _, s := range scoped {
		if hash[s] != true {
			hash[s] = true
		}
	}
	if (len(hash)) == len(scoped) {
		return true
	}
	return false
}

func main() {
	//count := 0

	// Open File
	f, err := os.Open("Real_signals.txt")
	if err != nil {
		panic(err)
	}
	bf := bufio.NewScanner(f)

	for bf.Scan() {
		signal := bf.Text()
		fmt.Println(signal)
		// just one scoped to 4 slices
		for i := 0; i <= len(signal)-14; i++ {
			marker := string(signal)
			scoped := marker[i : i+14]
			fmt.Println(scoped)
			fmt.Println(i+3, string(marker[i+13]))
			fmt.Println(checkfordup14(scoped))
			//if scoped is unique then break
			if checkfordup14(scoped) == true {
				fmt.Printf("Unique marker: %d \n", i+14)
				break
			}

		}

		defer f.Close()
	}
}
