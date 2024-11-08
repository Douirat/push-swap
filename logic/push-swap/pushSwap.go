package pushswap

import "fmt"

// Create a structure t hold both of my structs:
type Stacks struct {
	StackA []int
	StackB []int
}

// Instantiate the stacks struct:
func Init() *Stacks {
	stacks := &Stacks{}
	return stacks
}

// pa push the top first element of stack b to stack a:
func (stacks *Stacks) PB() {
	if len(stacks.StackA) == 0 {
		fmt.Println("The stack A is empty!!!")
		return
	}
	stacks.StackB = append(stacks.StackB, stacks.StackA[len(stacks.StackA)-1])
	stacks.StackA = stacks.StackA[:len(stacks.StackA)-1]
}

// pa push the top first element of stack b to stack a:
func (stacks *Stacks) PA() {
	if len(stacks.StackB) == 0 {
		fmt.Println("The stack A is empty!!!")
		return
	}
	stacks.StackA = append(stacks.StackA, stacks.StackB[len(stacks.StackB)-1])
	stacks.StackB = stacks.StackB[:len(stacks.StackB)-1]
}

// sa swap first 2 elements of stack a
func (stacks *Stacks) SA() {
	if len(stacks.StackA) < 2 {
		fmt.Println("Not enough values!!!")
		return
	}
	stacks.StackA[len(stacks.StackA)-1] += stacks.StackA[len(stacks.StackA)-2]
	stacks.StackA[len(stacks.StackA)-2] = stacks.StackA[len(stacks.StackA)-1] - stacks.StackA[len(stacks.StackA)-2]
	stacks.StackA[len(stacks.StackA)-1] -= stacks.StackA[len(stacks.StackA)-2]
}

// sb swap first 2 elements of stack b
func (stacks *Stacks) SB() {
	if len(stacks.StackB) < 2 {
		fmt.Println("Not enough values!!!")
		return
	}
	stacks.StackB[len(stacks.StackB)-1] += stacks.StackB[len(stacks.StackB)-2]
	stacks.StackB[len(stacks.StackB)-2] = stacks.StackB[len(stacks.StackB)-1] - stacks.StackB[len(stacks.StackB)-2]
	stacks.StackB[len(stacks.StackB)-1] -= stacks.StackB[len(stacks.StackB)-2]
}

// ss execute sa and sb:
func (stacks *Stacks) SS() {
	stacks.SA()
	stacks.SB()
}

// rra reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
func (stacks *Stacks) RRA() {
	temp := make([]int, len(stacks.StackA))
	copy(temp, stacks.StackA)
	for i := len(stacks.StackA) - 1; i >= 0; i-- {
		stacks.StackA[len(stacks.StackA)-i-1] = temp[i]
	}
}

// rra reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
func (stacks *Stacks) RRB() {
	temp := make([]int, len(stacks.StackB))
	copy(temp, stacks.StackB)
	for i := len(stacks.StackB) - 1; i >= 0; i-- {
		stacks.StackB[len(stacks.StackB)-i-1] = temp[i]
	}
}

// ra rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
func (stacks *Stacks) RA() {
	temp := make([]int, len(stacks.StackA))
	copy(temp, stacks.StackA)
	for i := len(stacks.StackA) - 1; i >= 0; i-- {
		ind := (i + 1) % (len(stacks.StackA))
		stacks.StackA[ind] = temp[i]
	}
}


// rb rotate stack b:
func (stacks *Stacks) RB() {
	temp := make([]int, len(stacks.StackB))
	copy(temp, stacks.StackB)
	for i := len(stacks.StackB) - 1; i >= 0; i-- {
		ind := (i + 1) % (len(stacks.StackB))
		stacks.StackB[ind] = temp[i]
	}
}

// rrr execute rra and rrb:
func (stacks *Stacks) RRR() {
	stacks.RRA()
	stacks.RRB()
}

// Sorting application:
func (stacks *Stacks) Sort() {
	stacks.PB()
	stacks.PB()
	stacks.PB()
	stacks.SA()
	stacks.RRA()
	// stacks.PA()
	// stacks.SA()
	// stacks.SB()
	// stacks.SS()
	// stacks.RRA()
	// stacks.RRR()
	// stacks.RB()
}
