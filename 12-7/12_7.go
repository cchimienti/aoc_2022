package main

/*
	Implementing a tree based on file system output, then getting and storing filesizes
	Calculate all dirs <

	To begin, find all of the directories with a total size of at most 100000, then calculate the sum of their total sizes.
	In the example above, these directories are a and e; the sum of their total sizes is 95437 (94853 + 584).
	(As in this example, this process can count files more than once!)

		Find all of the directories with a total size of at most 100000. What is the sum of the total sizes of those directories?
*/

import (
	dirs "12_7/pkg"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// var Root dirs.Dir()
var whereto string
var gotoDir dirs.Dir

func main() {
	// Open File
	f, err := os.Open("terminal.txt")
	if err != nil {
		panic(err)
	}
	bf := bufio.NewScanner(f)

	for bf.Scan() {
		dumpline := string(bf.Text())
		if string(dumpline[0]) == "$" {
			if string(dumpline[2:4]) == "cd" { // go either up, down in tree, or back to beginning w /
				fmt.Println(dumpline)
				whereto = string(dumpline[5:])
				fmt.Printf("whereto %s \n", whereto)
				if whereto == "/" { // Go back to root node
					fmt.Printf("here we go again \n")
					// Create root
					Root := dirs.Dir{Dir: "root"}
					gotoDir = Root
					fmt.Println(Root)
				} else if whereto == ".." { // Go back one node
					// GoToParent
				} else { // Go to that connected node
					//GoTo
				}
			} else if string(dumpline[2:4]) == "ls" {
				// need to pay attention until another command is called, but note value of whereto
				//fmt.Println(dumpline)
			}
		} else if string(dumpline[0:3]) == "dir" {
			_, dir, _ := strings.Cut(dumpline, " ") //note dir, don't assign it a value though, just child node of current whereto node
			//fmt.Println(dir)
			// from string to type Dir
			D := dirs.Dir{Dir: dir, Parent: &gotoDir}
			fmt.Println(D)
			//d := {}

		} else {
			size, file, _ := strings.Cut(dumpline, " ") // file and size, objects at end of tree
			fmt.Printf("%s = %s \n", file, size)
		}
		defer f.Close()
	}
}
