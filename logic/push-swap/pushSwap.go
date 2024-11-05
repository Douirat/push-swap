package pushswap

import "fmt"

// Create a list type tp hold the stack:
type Stack []int


// Create a structure to hold the tow stacks a and b:
type Stacks struct {
	A Stack
	B Stack
}

// Allocate a new stack:
func  AllocateStacks() *Stacks {
	return &Stacks{}
}


func PegeonSort(arr []int, n int) {
	fmt.Println(arr)
	Max := arr[0]
	Min := arr[0]
	index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < Min {
			Min = arr[i]
		} else if arr[i] > Max {
			Max = arr[i]
		}
	}
	fmt.Printf("The max is: %d and the min %d\n", Max, Min)
	Range := (Max - Min) + 1
	pigeonHoles := make([]int, Range)
	for j := 0; j < n; j++ {
		pigeonHoles[(arr[j]-Min)]++
	}
	fmt.Println(pigeonHoles)
	for k := 0; k < Range; k++ {
		for pigeonHoles[k] > 0 {
			arr[index] = k + Min
			pigeonHoles[k]--
			index++
		}
	}
	fmt.Println(arr)
}
