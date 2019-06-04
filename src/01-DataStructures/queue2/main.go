package main

import "fmt"

// START OMIT
type Queue struct {
	items chan int
}

func (q *Queue) Enqueue(v int) {
	q.items <- v
}

func (q *Queue) Dequeue() int {
	return <-q.items
}

// END OMIT

// START2 OMIT
func main() {
	q := Queue{
		items: make(chan int, 16),
	}

	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}

// END2 OMIT
