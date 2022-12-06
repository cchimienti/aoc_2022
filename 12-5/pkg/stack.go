package stack

// Crate stackin
// is empty
// push
// pop
// get top = pop

type Stack []string

func (st *Stack) IsEmpty() bool {
	return len(*st) == 0
}

//func RemoveIndex(s []string, index string) []string {
//	return append(s[:index], s[index+1:]...)
//}

func (st *Stack) Pop() (string, bool) {
	if st.IsEmpty() {
		return "", false
	} else {
		index := len(*st) - 1
		element := (*st)[index]
		//fmt.Printf("popped: %s \n", element)
		*st = (*st)[:index]

		return element, true
	}

}

func (st *Stack) Push(incoming string) {
	*st = append(*st, incoming)
}

func (st *Stack) Get() Stack {
	return st.Get()
}
