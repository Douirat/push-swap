package utils

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Douirat/push-swap/logic/push-swap"
)

// Split the input argument into a slice of strings:
func SplitInput(stacks *pushswap.Stacks, input string) error {
	visited := make(map[int]bool)
	slc := strings.Split(input, " ")
	if len(slc) == 0 {
		return errors.New("1")
	}
	for i := len(slc) - 1; i >= 0; i-- {
		dig, err := strconv.Atoi(slc[i])
		if err != nil {
			return err
		}
		if !visited[dig] {
			stacks.InsertIntoA(dig)
			visited[dig] = true
		} else {
			return errors.New("1")
		}
	}
	return nil
}
