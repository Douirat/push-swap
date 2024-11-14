package pushswap

import "fmt"

// A structr to represent the stack:
type Node struct {
	value       int
	Left, Right *Node
}

// Create a structure to hold both of my structs:
type Stacks struct {
	HeadStackA  *Node
	HeadStackkB *Node
}

// Instantiate a new stacks:
func Init() *Stacks {
	stacks := new(Stacks)
	stacks.HeadStackA = new(Node)
	stacks.HeadStackkB = new(Node)
	return stacks
}

// Instantiate a new node:
func NewNode(val int) *Node {
	node := new(Node)
	node.value = val
	node.Left = nil
	node.Right = nil
	return node
}

// Insert the user input to the stack a:
func (stacks *Stacks) Insert(value int) {
	node := NewNode(value)
	if stacks.HeadStackA == nil {
		stacks.HeadStackA = node
	} else {
		stacks.HeadStackA.Right = node
		stacks.HeadStackA.Right.Left = stacks.HeadStackA
		stacks.HeadStackA = node
	}
}

// Display stack:
func (stacks *Stacks) Display() {
	Temp := stacks.HeadStackA
	for {
		fmt.Println(Temp.value)
		Temp = Temp.Left
		if Temp.Left == nil {
			break
		}
	}
}

// 
