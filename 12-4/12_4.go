package main

/*


.234.....  2-4
.....678.  6-8

.23......  2-3
...45....  4-5

....567..  5-7
......789  7-9

.2345678.  2-8
..34567..  3-7

.....6...  6-6
...456...  4-6

.23456...  2-6
...45678.  4-8
____________________

	Given two sections x-y, x1-y1, find how many occurences of total overlap happen
	In my head I picture these as two venn diagrams & need to count how many are concentric circles

	At first glance I feel like I'll need slices
	Create a list of length - largest num given in the line;

	there has to be a way to do this easily with matrices for sure

*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ew don't look at this
func quick_conv(starts, ends string, err bool) (start, end int) {
	start, _ = strconv.Atoi(starts)
	end, _ = strconv.Atoi(ends)
	return start, end
}

var tally int

func main() {
	// Open File
	f, err := os.Open("Real_sections.txt")
	if err != nil {
		panic(err)
	}
	bf := bufio.NewScanner(f)

	for bf.Scan() {
		// Cut up into sections, idk why i named these weird
		line1, line2, _ := strings.Cut(bf.Text(), ",")
		s1, e1 := quick_conv(strings.Cut(line1, "-"))
		s2, e2 := quick_conv(strings.Cut(line2, "-"))
		fmt.Println(s1, e1)
		fmt.Println(s2, e2)

		// Check if Overlaps
		if e1 >= s2 && s1 <= e2 {
			tally++
		}

		/*
			if (s1 < s2) && (e1 < e2) {
				//fmt.Printf("No")
				//Nope
			} else if (s1 > s2) && (e1 > e2) {
				//fmt.Printf("No")
				// No
			} else if (s1 <= s2) && (e1 >= e2) {
				fmt.Printf("Yes")
				tally++
				// Yes
			} else if (s1 >= s2) && (e1 <= e2) {
				fmt.Printf("Yes")
				tally++
				// Yes
			}
		*/
		fmt.Printf("Total: %d\n", tally)

	}

	defer f.Close()

}
