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
	for _, v := range slc {
		dig, err := strconv.Atoi(v)
		if err != nil {
			return
		}
		stack.StackA = append(stack.StackA, dig)
	}
}
