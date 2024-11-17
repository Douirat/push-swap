package pushswap

import "fmt"

// Create a structureto represent a node in a circular doubly linked list:
type Node struct {
	Value          int
	Previous, Next *Node
}

// A structr to represent the stack:
type Stack struct {
	Values []int
}

// Create a structure to hold both of my structs:
type Stacks struct {
	HeadStackA  *Node
	HeadStackkB *Stack
}

// Instantiate a new stacks:
func Init() *Stacks {
	stacks := new(Stacks)
	return stacks
}

// Instantiate a new node:
func NewNode(val int) *Node {
	node := new(Node)
	node.Value = val
	node.Previous = nil
	node.Next = nil
	return node
}

// Insert the user input to the stack a:
func (stacks *Stacks) Initialize(value int) {
	node := NewNode(value)
	if stacks.HeadStackA == nil {
		node.Previous = node
		node.Next = node
		stacks.HeadStackA = node
	} else {
		node.Previous = stacks.HeadStackA
		node.Next = node.Previous.Next
		stacks.HeadStackA.Next = node
		node.Next.Previous = node
		stacks.HeadStackA = node
	}
}

// Display the content of my stack;
func (stacks *Stacks) Display() {
	if stacks.HeadStackA == nil {
		return
	}
	Temp := stacks.HeadStackA
	for {
		fmt.Println(Temp.Value)
		if Temp.Previous == stacks.HeadStackA {
			break
		}
		Temp = Temp.Previous
	}
}

// Display stack B:
func (stacks *Stacks) PrintStack() {
	if len(stacks.HeadStackkB.Values) == 0 {
		fmt.Println("The stack is empty")
		return
	}
	fmt.Println("The Values of stack B are:")
	for i := len(stacks.HeadStackkB.Values) - 1; i >= 0; i-- {
		fmt.Println(stacks.HeadStackkB.Values[i])
	}
}

// Check if my stack is sorted:
func (stacks *Stacks) IsSorted() bool {
	if stacks.HeadStackA == nil {
		return false
	}
	Temp := stacks.HeadStackA
	for {
		if Temp.Previous == stacks.HeadStackA {
			return true
		}
		if Temp.Value > Temp.Previous.Value {
			return false
		}
		Temp = Temp.Previous
	}
}

// pa push the top first element of stack b to stack a:
func (stacks *Stacks) PushAtoB() {
	if stacks.HeadStackA == nil {
		return
	}
	stacks.HeadStackkB.Values = append(stacks.HeadStackkB.Values, stacks.HeadStackA.Value)
	Prev := stacks.HeadStackA.Previous
	Prev.Next = stacks.HeadStackA.Next
	Prev.Next.Previous = Prev
	stacks.HeadStackA = Prev
	if stacks.HeadStackA == stacks.HeadStackA.Next && stacks.HeadStackA == stacks.HeadStackA.Previous {
		stacks.HeadStackA = nil
	}
}

// pb push the top first element of stack a to stack b:
func (stacks *Stacks) PushBtoA() {
	if len(stacks.HeadStackkB.Values) == 0 {
		return
	}
	stacks.Initialize(stacks.HeadStackkB.Values[len(stacks.HeadStackkB.Values)-1])
	stacks.HeadStackkB.Values = stacks.HeadStackkB.Values[:len(stacks.HeadStackkB.Values)-1]
}

// sa swap first 2 elements of stack a:
func (stacks *Stacks) SwapA() {
	if stacks.HeadStackA == nil || stacks.HeadStackA.Previous == nil {
		return
	}
	stacks.HeadStackA.Value = stacks.HeadStackA.Value + stacks.HeadStackA.Previous.Value
	stacks.HeadStackA.Previous.Value = stacks.HeadStackA.Value - stacks.HeadStackA.Previous.Value
	stacks.HeadStackA.Value = stacks.HeadStackA.Value - stacks.HeadStackA.Previous.Value
}

// sb swap first 2 elements of stack b:
func (stacks *Stacks) SwapB() {
	if len(stacks.HeadStackkB.Values) < 2 {
		return
	}
	stacks.HeadStackkB.Values[len(stacks.HeadStackkB.Values)-1] = stacks.HeadStackkB.Values[len(stacks.HeadStackkB.Values)-1] + stacks.HeadStackkB.Values[len(stacks.HeadStackkB.Values)-2]
	stacks.HeadStackkB.Values[len(stacks.HeadStackkB.Values)-2] = stacks.HeadStackkB.Values[len(stacks.HeadStackkB.Values)-1] - stacks.HeadStackkB.Values[len(stacks.HeadStackkB.Values)-2]
	stacks.HeadStackkB.Values[len(stacks.HeadStackkB.Values)-1] = stacks.HeadStackkB.Values[len(stacks.HeadStackkB.Values)-1] - stacks.HeadStackkB.Values[len(stacks.HeadStackkB.Values)-2]
}

// The sorter function:
func (stacks *Stacks) Sort() {
	stacks.PushAtoB()
	// stacks.PushBtoA()
	stacks.PushAtoB()
	stacks.PushAtoB()
	// stacks.PushAtoB()
	// stacks.SwapA()
	stacks.SwapB()
	stacks.SwapB()
}
