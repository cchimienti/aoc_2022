package ui

import (
	"12_9/src/head"
	"time"

	"github.com/gdamore/tcell"
)

// Draw out the Head, Tail map

type Display struct {
	Screen tcell.Screen
	Head   head.Head
}

func (d *Display) Run() {

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	d.Screen.SetStyle(defStyle)
	//width, height := d.Screen.Size()
	ropeStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)

	//for {
	d.Screen.Clear()
	//	d.Head.Update(width, height)
	d.Screen.SetContent(d.Head.X, d.Head.Y, ' ', nil, ropeStyle)
	time.Sleep(40 * time.Millisecond)
	d.Screen.Show()
	//}
}
