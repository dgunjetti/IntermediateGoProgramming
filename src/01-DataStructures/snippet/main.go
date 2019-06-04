package main

// START OMIT
type Element struct {
	next, prev *Element
	Value      interface{}
}

// END OMIT

// START1 OMIT
type List struct {
	root Element // sentinel list element, only &root, root.prev, and root.next are used
	len  int     // current list length excluding (this) sentinel element
}

// END1 OMIT

// START2 OMIT
func New() *List

// END2 OMIT

// START3 OMIT
func (l *List) PushFront(v interface{}) *Element

// END3 OMIT

// START4 OMIT
func (l *List) PushBack(v interface{}) *Element

// END4 OMIT


// START5 OMIT
for e := l.Front(); e != nil; e = e.Next() {
	// do something with e.Value
}
// END5 OMIT

// START6 OMIT
func (l *List) Remove(e *Element) interface{}
// END6 OMIT

func main() {

}
