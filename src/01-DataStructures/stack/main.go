package main

import "fmt"

//START OMIT
type Stack struct {
	items []int
}

func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	item, items := s.items[len(s.items)-1], s.items[:len(s.items)-1]
	s.items = items
	return item
}

//END OMIT

//START2 OMIT
func main() {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

//END2 OMIT
