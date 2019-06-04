package main

import "fmt"

// START OMIT
type Queue struct {
	items []int
}

func (q *Queue) Enqueue(v int) {
	q.items = append(q.items, v)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	item, items := q.items[0], q.items[1:len(q.items)]
	q.items = items
	return item
}

// END OMIT

// START2 OMIT
func main() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}

// END2 OMIT
