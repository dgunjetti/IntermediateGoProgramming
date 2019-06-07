package main

import "fmt"

//START1 OMIT
type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}

//END1 OMIT
func main() {
	//START2 OMIT
	root := Tree{
		LeafValue: 0,
		Right: &Tree{
			LeafValue: 5,
			Right:     &Tree{6, nil, nil},
			Left:      nil,
		},
		Left: &Tree{4, nil, nil},
	}
	//END2 OMIT

	fmt.Printf("%#v\n", root)
}
