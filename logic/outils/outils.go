package outils

import (
	"strconv"
	"strings"
	p "github.com/Douirat/push-swap/logic/push-swap"
)

// Split the input argument into a slice of strings:
func SplitInput(stack *p.Stacks, input string)  {
	slc := strings.Split(input, " ")
	if len(slc) == 0 {
		return
	}
	for i:=len(slc)-1; i>=0; i--{
		dig, err := strconv.Atoi(slc[i])
		if err != nil {
			return
		}
		stack.StackA = append(stack.StackA, dig)
	}
}
