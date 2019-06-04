package main

import (
	"container/list"
	"fmt"
)

//START OMIT
func iter(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func remove(v int, l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		if v == e.Value.(int) {
			l.Remove(e)
		}
	}
}

//END OMIT

func main() {
	// START1 OMIT
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	// END1 OMIT
	l.PushBack(3)

	iter(l)
	remove(2, l)
	iter(l)
}
