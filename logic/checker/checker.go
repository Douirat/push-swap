package checker

import (
	"fmt"

	pushswap "github.com/Douirat/push-swap/logic/push-swap"
)

// Create a checker function:
func Check(stacks *pushswap.Stacks, input []string) bool {
	fmt.Println(stacks.HeadStackA)
	stacks.Operations = input
	if len(stacks.Operations) == 0 {
		return false
	}
	for _, Operation := range stacks.Operations {
		switch Operation {
		case "pa":
			stacks.PushBtoA()
		case "pb":
			stacks.PushAtoB()
		case "sa":
			stacks.SwapA()
		case "sb":
			stacks.SwapB()
		case "ss":
			stacks.SwapA()
			stacks.SwapB()
		case "ra":
			stacks.RotateStackA()
		case "rb":
			stacks.RotateStackB()
		case "rr":
			stacks.RotateStackAandB()
		case "rra":
			stacks.ReverseRotateA()
		case "rrb":
			stacks.ReverseRotateB()
		case "rrr":
			stacks.ReverseRotateAandB()
		default:
			return false
		}
	}

	stacks.DisplayStackA()
	stacks.DisplayOperations()
	return stacks.IsSorted()
}