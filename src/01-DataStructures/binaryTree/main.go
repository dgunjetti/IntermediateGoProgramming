package main

import "fmt"

// START1 OMIT
type Tree struct {
	node *Node
}

func (t *Tree) Insert(v int) *Tree {
	if t.node == nil {
		t.node = &Node{value: v}
	} else {
		t.node.insert(v)
	}
	return t
}

// END1 OMIT

//START2 OMIT
type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) insert(v int) {
	if n.value <= v {
		if n.left == nil {
			n.left = &Node{value: v}
		} else {
			n.left.insert(v)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: v}
		} else {
			n.right.insert(v)
		}
	}
}

//END2 OMIT
//START3 OMIT
func (n *Node) isExist(v int) bool {
	if n == nil {
		return false
	}
	if n.value == v {
		return true
	}
	if n.value < v {
		return n.left.isExist(v)
	}

	if n.value > v {
		return n.right.isExist(v)
	}
	return false
}

//END3 OMIT

//START4 OMIT
func PrintNode(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.value)
	PrintNode(n.left)
	PrintNode(n.right)
}

func main() {
	t := &Tree{}
	t.Insert(10).Insert(5).Insert(20).Insert(0) // insert is return tree.
	PrintNode(t.node)
	fmt.Println(t.node.isExist(20))
	fmt.Println(t.node.isExist(9))
}

//END4 OMIT
