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
		ind := (i - 1 + len(stacks.StackA)) % (len(stacks.StackA))
		stacks.StackA[ind] = temp[i]
	}
}

// rra reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
func (stacks *Stacks) RRB() {
	temp := make([]int, len(stacks.StackB))
	copy(temp, stacks.StackB)

	for i := 0; i < len(stacks.StackB); i++ {
		// Calculate the new index with wrap-around
		ind := (i - 1 + len(stacks.StackB)) % len(stacks.StackB)
		stacks.StackB[ind] = temp[i]
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
	Min, Max := stacks.StackA[0], stacks.StackA[0]
	index := 0
	for i := 0; i < len(stacks.StackA); i++ {
		if stacks.StackA[i] < Min {
			Min = stacks.StackA[i]
		} else if stacks.StackA[i] > Max {
			Max = stacks.StackA[i]
		}
	}
	fmt.Printf("The Max is :%d and min is %d\n", Max, Min)
	Range := Max - Min + 1
	Track := make([]int, Range)
	for j := 0; j < len(stacks.StackA); j++ {
		Track[(stacks.StackA[j]-Min)]++
	}
	Temp := make([]int, len(stacks.StackA))
	for x := Range - 1; x >= 0; x-- {
		if Track[x] > 1 {
			return 
		}
		if Track[x] != 0{
			Temp[index] = x + Min
			index++
		}
	}	
	fmt.Println(Track)
	fmt.Println(Temp)
}
