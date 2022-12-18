package monkeys

// monkeys

type Operation func([]string, int) int

type MonkeyHoldings struct {
	Name        int
	WorryLevels []int
}

type Monkey struct {
	Name       int
	Starting   []int
	Operation  Operation
	Test       int
	Test_True  int //*Monkey.Name
	Test_False int //*Monkey.Name
	Save_Func  []string
}
