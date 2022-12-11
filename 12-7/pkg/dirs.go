package dirs

// Treeeees
// This'll be a neat lil filesystem

type File struct { //Edge, Valued
	Size   int
	File   string
	Parent *Dir
}

//type Root struct {
//	Name  string
//	Child []*Dir
//}

type Dir struct { //Node, UnValued
	Child []*File
	//Size   int // this is default 0 by Go rules
	Parent *Dir
	Dir    string
}

func (d *Dir) InsertDir(parent *Dir, current Dir, dirname string) {
	d.Parent = parent
	d.Dir = dirname
}
func (f *File) InsertFile(current *Dir, file *File, filename string, size int) {
	current.SetChild(file)
	f.Parent = current
	f.File = filename
	f.Size = size
}

func (d *Dir) SetChild(file *File) {
	d.Child = append(d.Child, file)
}

/*
func (d *Dir) GoBack(current Dir) {
	parent := d.GetParent()
	d.GoTo(parent)
}

func (d *Dir) GoTo(new Dir) { //Traverse
	new.
		d.Dir = new
}

func (d *Dir) GetCurrentDir() {
	return d.Dir
}

func (d *Dir) LS() {
	return d.Child
}

func (f *File) GetSize() int {
	return f.Size
}

func (f *File) GetParent() *Dir {
	return &f.Parent
}

func (d *Dir) GetParent() *Dir {
	return &d.Parent
}
*/
