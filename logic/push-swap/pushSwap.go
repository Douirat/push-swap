package pushswap

import "fmt"

// Create a structureto represent a node in a circular doubly linked list:
type Node struct {
	Value          int
	Previous, Next *Node
}

// Create a structure to hold both of my structs:
type Stacks struct {
	HeadStackA *Node
	HeadStackB *Node
	Operations []string
	Size       int
}

// Instantiate a new stacks:
func Init() *Stacks {
	stacks := new(Stacks)
	stacks.Size = 0
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
func (stacks *Stacks) InsertIntoA(value int) {
	node := NewNode(value)
	if stacks.HeadStackA == nil {
		node.Previous = node
		node.Next = node
		stacks.HeadStackA = node
		stacks.Size++
	} else {
		node.Previous = stacks.HeadStackA
		node.Next = node.Previous.Next
		stacks.HeadStackA.Next = node
		node.Next.Previous = node
		stacks.HeadStackA = node
		stacks.Size++
	}
}

// Insert a new value to stack B:
func (stacks *Stacks) InsertIntoB(value int) {
	node := NewNode(value)
	if stacks.HeadStackB == nil {
		node.Previous = node
		node.Next = node
		stacks.HeadStackB = node
	} else {
		node.Previous = stacks.HeadStackB
		node.Next = node.Previous.Next
		stacks.HeadStackB.Next = node
		node.Next.Previous = node
		stacks.HeadStackB = node
	}
}

// Display the content of my stack;
func (stacks *Stacks) DisplayStackA() {
	if stacks.HeadStackA == nil {
		fmt.Println("The stack A is empty!")
		return
	}
	fmt.Println("The Values of stack A are:")
	Temp := stacks.HeadStackA
	for {
		fmt.Println(Temp.Value)
		if Temp.Previous == stacks.HeadStackA {
			break
		}
		Temp = Temp.Previous
	}
}

// Display the content of my stack;
func (stacks *Stacks) DisplayStackB() {
	if stacks.HeadStackB == nil {
		fmt.Println("The stack B is empty!")
		return
	}
	fmt.Println("The Values of stack B are:")
	Temp := stacks.HeadStackB
	for {
		fmt.Println(Temp.Value)
		if Temp.Previous == stacks.HeadStackB {
			break
		}
		Temp = Temp.Previous
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
			return Temp.Value >= Temp.Next.Value
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
	if stacks.HeadStackA == stacks.HeadStackA.Next && stacks.HeadStackA == stacks.HeadStackA.Previous {
		stacks.InsertIntoB(stacks.HeadStackA.Value)
		stacks.HeadStackA = nil
		return
	}
	stacks.InsertIntoB(stacks.HeadStackA.Value)
	Prev := stacks.HeadStackA.Previous
	Prev.Next = stacks.HeadStackA.Next
	Prev.Next.Previous = Prev
	stacks.HeadStackA = Prev
}

// pa push the top first element of stack b to stack a:
func (stacks *Stacks) PushBtoA() {
	if stacks.HeadStackB == nil {
		return
	}
	if stacks.HeadStackB == stacks.HeadStackB.Next && stacks.HeadStackB == stacks.HeadStackB.Previous {
		stacks.InsertIntoA(stacks.HeadStackB.Value)
		stacks.HeadStackB = nil
		return
	}
	stacks.InsertIntoA(stacks.HeadStackB.Value)
	Prev := stacks.HeadStackB.Previous
	Prev.Next = stacks.HeadStackB.Next
	Prev.Next.Previous = Prev
	stacks.HeadStackB = Prev
}

// sa swap first 2 elements of stack A:
func (stacks *Stacks) SwapA() {
	if stacks.HeadStackA == nil || stacks.HeadStackA.Previous == nil {
		return
	}
	stacks.HeadStackA.Value = stacks.HeadStackA.Value + stacks.HeadStackA.Previous.Value
	stacks.HeadStackA.Previous.Value = stacks.HeadStackA.Value - stacks.HeadStackA.Previous.Value
	stacks.HeadStackA.Value = stacks.HeadStackA.Value - stacks.HeadStackA.Previous.Value
}

// sa swap first 2 elements of stack B:
func (stacks *Stacks) SwapB() {
	if stacks.HeadStackB == nil || stacks.HeadStackB.Previous == nil {
		return
	}
	stacks.HeadStackB.Value = stacks.HeadStackB.Value + stacks.HeadStackB.Previous.Value
	stacks.HeadStackB.Previous.Value = stacks.HeadStackB.Value - stacks.HeadStackB.Previous.Value
	stacks.HeadStackB.Value = stacks.HeadStackB.Value - stacks.HeadStackB.Previous.Value
}

// ra rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
func (stacks *Stacks) RotateStackA() {
	if stacks.HeadStackA == nil {
		return
	}
	stacks.HeadStackA = stacks.HeadStackA.Previous
}

// rb rotate stack b (shift up all elements of stack a by 1, the first element becomes the last one)
func (stacks *Stacks) RotateStackB() {
	if stacks.HeadStackB == nil {
		return
	}
	stacks.HeadStackB = stacks.HeadStackB.Previous
}

// Rotate A and B:
func (stacks *Stacks) RotateStackAandB() {
	stacks.RotateStackA()
	stacks.RotateStackB()
}

// rra reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
func (stacks *Stacks) ReverseRotateA() {
	if stacks.HeadStackA == nil {
		return
	}
	stacks.HeadStackA = stacks.HeadStackA.Next
}

// Reverse rotate B:
func (stacks *Stacks) ReverseRotateB() {
	if stacks.HeadStackB == nil {
		return
	}
	stacks.HeadStackB = stacks.HeadStackB.Next
}

// Reverse rotate A and reverse rotate B:
func (stacks *Stacks) ReverseRotateAandB() {
	stacks.ReverseRotateA()
	stacks.ReverseRotateB()
}

// Extract Minimum value from the stack:
func (stacks *Stacks) ExtractMin() int {
	Min := stacks.HeadStackA.Value
	Temp := stacks.HeadStackA
	if Temp.Previous.Value < Temp.Value {
		Min = Temp.Previous.Value
	}
	Temp = Temp.Previous
	for Temp != stacks.HeadStackA {
		if Temp.Value < Min {
			Min = Temp.Value
		}
		Temp = Temp.Previous
	}
	return Min
}

// The sorter function: ./checker "0 one 2 3"
func (stacks *Stacks) Sort() {
	if stacks.HeadStackA == nil || stacks.IsSorted() {
		return
	}
	for stacks.Size > 3 {
		if stacks.HeadStackA.Value == stacks.ExtractMin() {
			stacks.PushAtoB()
			stacks.Size--
			stacks.Operations = append(stacks.Operations, "pb")
			continue
		}
		if stacks.HeadStackA.Previous.Value == stacks.ExtractMin() {
			stacks.SwapA()
			stacks.PushAtoB()
			stacks.Size--
			stacks.Operations = append(stacks.Operations, "sa")
			stacks.Operations = append(stacks.Operations, "pb")
		} else {
			stacks.RotateStackA()
			stacks.Operations = append(stacks.Operations, "ra")
		}
	}
	if !stacks.IsSorted() {
		switch {
		case (stacks.HeadStackA.Value == stacks.ExtractMin() && stacks.HeadStackA.Previous.Value > stacks.HeadStackA.Next.Value):
			stacks.SwapA()
			stacks.RotateStackA()
			stacks.Operations = append(stacks.Operations, "sa")
			stacks.Operations = append(stacks.Operations, "ra")
		case (stacks.HeadStackA.Previous.Value == stacks.ExtractMin() && stacks.HeadStackA.Next.Value > stacks.HeadStackA.Value):
			stacks.SwapA()
			stacks.Operations = append(stacks.Operations, "sa")
		case (stacks.HeadStackA.Next.Value == stacks.ExtractMin() && stacks.HeadStackA.Previous.Value < stacks.HeadStackA.Value):
			stacks.SwapA()
			stacks.ReverseRotateA()
			stacks.Operations = append(stacks.Operations, "sa")
			stacks.Operations = append(stacks.Operations, "rra")

		case (stacks.HeadStackA.Previous.Value == stacks.ExtractMin() && stacks.HeadStackA.Next.Value < stacks.HeadStackA.Value):
			stacks.RotateStackA()
			stacks.Operations = append(stacks.Operations, "ra")
		}
	}
	for stacks.HeadStackB != nil {
		stacks.PushBtoA()
		stacks.Operations = append(stacks.Operations, "pa")
	}
}

// Display operations:
func (stacks *Stacks) DisplayOperations() {
	for i := 0; i < len(stacks.Operations); i++ {
		fmt.Println(stacks.Operations[i])
	}
}
