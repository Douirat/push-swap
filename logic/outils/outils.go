package outils

import (
	"strconv"
	"strings"
)

// Split the input argument into a slice of strings:
func SplitInput(input string) []int {
	slc := strings.Split(input, " ")
	ints := []int{}
	if len(slc) == 0 {
		return nil
	}
	for _, v := range slc {
		dig, err := strconv.Atoi(v)
		if err != nil {
			return nil
		}
		ints = append(ints, dig)
	}
	return ints
}
