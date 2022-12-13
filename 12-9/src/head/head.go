package head

type Head struct {
	X int
	Y int
	//Xspeed int
	//Yspeed int
}

type Tail struct {
	X int
	Y int
}

//func (h *Head) ChangeDir(vertical int, horizontal int) {
//	h.Yspeed = vertical
//	h.Xspeed = horizontal
//}

func (h *Head) Update(x int, y int) {
	h.Y += y
	h.X += x
}
