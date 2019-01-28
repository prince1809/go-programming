package intlist

type IntList struct {
	Value int
	Tails *IntList
}

func (list *IntList) sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tails.sum()
}
