package main

import (
	monkeys "12_11/pkg"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Monkeys??
To get your stuff back, you need to be able to predict where the monkeys will throw your items.
After some careful observation, you realize the monkeys operate based on how worried you are about each item. v relatable

Each monkey has several attributes:

Starting items lists your -worry level- for each item the monkey is currently holding in the order they will be inspected.
Operation shows how your -worry level changes- as that monkey inspects an item. (An operation like new = old * 5 means that your
worry level after the monkey inspected the item is five times whatever your worry level was before inspection.)
Test shows how the monkey -uses your worry level to decide where to throw an item next-.
If true shows what happens with an item if the Test was true.
If false shows what happens with an item if the Test was false.
*/
var items []int
var monkeyname int
var test int
var testt int
var testf int
var worry int
var constant int
var result monkeys.Monkey
var opp_func []string
var rounds int
var determine int

var Holdings = map[int]monkeys.MonkeyHoldings{} //Maybe update this into pointer later
var SadCage = map[int]monkeys.Monkey{}

func RemoveHolding() {
	mon, ok := Holdings[monkeyname]
	fmt.Printf("Removing from Monkey %d \n", monkeyname)
	if ok {
		if len(mon.WorryLevels) > 0 {
			mon.WorryLevels = mon.WorryLevels[1:]
			Holdings[monkeyname] = mon
		}
	}
}

// For each monkey holding an item, list what they're now holding
func AddHolding(monkey int, item int) {
	mon, ok := Holdings[monkey] //check for holdings
	//fmt.Println(mon)
	if ok {
		//fmt.Printf("Monkey %d Holding exists! \n", monkey)
		mon.WorryLevels = append(mon.WorryLevels, item)
		fmt.Println(mon)
		Holdings[monkey] = mon
	} else { // otherwise build it
		build := monkeys.MonkeyHoldings{
			Name:        monkey,
			WorryLevels: []int{item},
		}
		Holdings[monkey] = build //store in Monkey library
		fmt.Println(build)
	}
}

// Where does the Item Go
func MonkeyHolds(determine int) {
	worry_less := determine / 3
	fmt.Printf("( %d / 3 ) = %d ", determine, worry_less)
	fmt.Printf("/ %d = %d --> ", result.Test, worry_less%result.Test)
	if worry_less%result.Test == 0 {
		// Item will go to Test True Monkey
		fmt.Println(result.Test_True, worry_less)
		//if _, ok := Holdings[monkey]
		AddHolding(result.Test_True, worry_less)

	} else {
		// item will go to Test False Monkey
		fmt.Println(result.Test_False, worry_less)
		AddHolding(result.Test_False, worry_less)

	}
	//fmt.Println(result.Operation(items, opp_func))
}

func BuildAMonkey() {
	// Detect ops
	ops := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	//fmt.Println(items)
	mon, ok := Holdings[monkeyname] //check for holdings, add if some already calculated
	//fmt.Println(ok, mon)
	if ok {
		for _, v := range mon.WorryLevels {
			items = append(items, v)
			//fmt.Println(items)
		}
	}

	for _, v := range items {
		fmt.Printf("Inspecting item: %d \n", v)
		// remove previous worrylevel from items
		if len(items) > 0 {
			//fmt.Println(x)
			items = items[1:]
			//fmt.Println(items)
		}
		// Build a Monkey if not already build
		monk, ok := SadCage[monkeyname]
		if ok {
			result = monk
			fmt.Printf("~ Old Monkey ~ \n")
			result.Starting = items //update items
			//update oldworrylevel
			determine = result.Operation(result.Save_Func, v)

		} else {
			fmt.Printf("~ New Monkey ~\n")
			result = monkeys.Monkey{
				Name:     monkeyname,
				Starting: items,
				Operation: func(opp_func []string, old int) int {
					reg := opp_func[0]
					operator := opp_func[1]
					constt := opp_func[2]
					if reg == "old" {
						worry = old
					} else {
						worry, _ = strconv.Atoi(reg)
					}
					if constt == "old" {
						constant = old
					} else {
						constant, _ = strconv.Atoi(constt)
					}

					newWorryLevel := ops[operator](worry, constant)
					//fmt.Println(worry, operator, constant)

					return newWorryLevel
				},
				Test:       test,
				Test_True:  testt,
				Test_False: testf,
				Save_Func:  opp_func,
			}
			determine = result.Operation(result.Save_Func, v)
			//fmt.Println(result)
			SadCage[monkeyname] = result
			//fmt.Printf("SadCage: \n")
			//fmt.Println(SadCage[monkeyname])

		}
		//fmt.Println(determine)
		MonkeyHolds(determine) //Call for every Item encountered
		RemoveHolding()        //Remove worry level looked at
	}
}

func NotherRound(times int) {
	fmt.Println(SadCage)
	fmt.Println(Holdings)
	fmt.Printf("\n\n\n")
	fmt.Println(times + 1)
	fmt.Printf("\n\n\n")
	for _, v := range SadCage {

		monkeyname = v.Name
		v.Starting = nil
		fmt.Println(v)
		//items = nil
		BuildAMonkey()
	}
}

func main() {
	// Open File
	f, err := os.Open("notes.txt")
	if err != nil {
		panic(err)
	}
	bf := bufio.NewScanner(f)

	for bf.Scan() {
		// Read in input
		x := string(bf.Text())
		//fmt.Println(x)
		if len(x) <= 0 {
			BuildAMonkey()
			//clear []s But still need to reuse some calculations
			items = nil
			continue
		}
		if string(x[0]) == "M" {
			_, strmonkeyname, _ := strings.Cut(x, " ")
			monkeyname, _ = strconv.Atoi(string(strmonkeyname[0]))
			//BuildAHolding(monkeyname)
			//fmt.Println(monkeyname)
		} else if string(x[2]) == "S" {
			_, startingitems, _ := strings.Cut(x, ": ")
			// Jam into list
			sitems := strings.Split(startingitems, ", ")
			for _, v := range sitems {
				intitems, _ := strconv.Atoi(v)
				items = append(items, int(intitems))
			}
			//fmt.Println(items)
		} else if string(x[2]) == "O" {
			_, opp, _ := strings.Cut(x, "= ")
			opp_func = strings.Split(opp, " ")
			//fmt.Println(opp_func)
		} else if string(x[2]) == "T" {
			_, test_s, _ := strings.Cut(x, "by ")
			test, _ = strconv.Atoi(test_s)
			//fmt.Println(test)
		} else if string(x[7]) == "t" {
			_, testt_s, _ := strings.Cut(x, "monkey ")
			testt, _ = strconv.Atoi(testt_s)
			//fmt.Println(testt)
		} else if string(x[7]) == "f" {
			_, testf_s, _ := strings.Cut(x, "monkey ")
			testf, _ = strconv.Atoi(testf_s)
			//fmt.Println(testf)
		}

	}

	rounds = 19
	//clear items list
	for i := 0; i < rounds; i++ {
		NotherRound(i) //1
		//NotherRound() //2
		//NotherRound() //3

	}

	fmt.Println(Holdings)
	//fmt.Println(SadCage)
	defer f.Close()

}
