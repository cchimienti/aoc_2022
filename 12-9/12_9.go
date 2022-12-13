package main

/*
	Rope - Kinda looks like a smaller version of snake
	Have Tail T follow Head H;
*/

import (
	"12_9/src/head"
	"12_9/src/ui"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gdamore/tcell"
)

func displayed(dir string, steps int) {

	// Create new Display
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)

	// Set Head and Display Content
	head := head.Head{
		X: 5,
		Y: 10,
		//Xspeed: 0,
		//Yspeed: 0,
	}

	display := ui.Display{
		Screen: s,
		Head:   head,
	}
	go display.Run()

	for {
		switch event := display.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			display.Screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				display.Screen.Fini()
				os.Exit(0)
			} else if dir == "L" { //
				display.Head.Update(-steps, 0)
			} else if dir == "R" {
				display.Head.Update(steps, 0)
			} else if dir == "D" {
				display.Head.Update(0, -steps)
			} else if dir == "U" {
				display.Head.Update(0, steps)
			}
		}
	}

}

func main() {
	// Open File
	f, err := os.Open("motions.txt")
	if err != nil {
		panic(err)
	}
	bf := bufio.NewScanner(f)

	for bf.Scan() {
		instruct := string(bf.Text())
		dir, ssteps, _ := strings.Cut(instruct, " ")
		//fmt.Println(dir, ssteps)
		steps, _ := strconv.Atoi(ssteps)
		displayed(dir, steps)

		// Clear screen
		//s.Clear()

	}

	defer f.Close()

}
