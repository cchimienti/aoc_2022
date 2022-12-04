// main.go

//Ew just realzed this is gonna be a sort algo thing
// mwhatever, lets make it efficient
//technically a snack will never be 0 cal, right? -- no that's the point!!

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// stolen func cuz I got tired
func max(numbers map[int]int) (int, int, int) {
	var maxNumber, second, third int
	for _, k := range numbers {
		maxNumber = k
		break
	}
	for _, k := range numbers {
		if k > maxNumber {
			third = second
			second = maxNumber
			maxNumber = k
		} else if k > second {
			third = second
			second = k
		} else if k > third {
			third = k
		}
	}

	return maxNumber, second, third
}

func main() {
	//var beefiest int
	tallies := make(map[int]int)

	// chunk method
	f, err := os.Open("Real_Calories.txt")
	if err != nil {
		panic(err)
	}
	// bufio method (still need f)
	bf := bufio.NewScanner(f)
	//bf.Split(bufio.ScanWords)
	//fmt.Printf("5 bytes: %s\n", string(bytes_read))

	// --> first tally, then store, and either pull or sort
	// ugh just put tallys into an array, and splice it by largest in array
	//Scan line by line
	sum := 0
	little_fellas := 0
	for bf.Scan() {
		new_snack := bf.Text()
		if len(new_snack) > 0 {
			new_snack, _ := strconv.Atoi(bf.Text())
			sum += new_snack
			fmt.Printf("cal: %d, total_cal: %d \n", new_snack, sum)

		} else {
			//if sum > beefiest {
			//	fmt.Printf("norm beef %d", beefiest)
			//	beefiest := sum
			//	fmt.Printf("BEEF: %d \n", beefiest)
			//}
			little_fellas++
			tallies[little_fellas] = sum
			fmt.Println(tallies)
			fmt.Printf("little fella %d has %d cals \n", little_fellas, sum)
			sum = 0 //reset tally

		}

	}

	final, second, third := max(tallies)
	fmt.Printf("1st %d, 2nd %d 3rd %d  - tots = %d \n", final, second, third, (final + second + third))

	defer f.Close()

}
