package outils

import(
	"strings"
)

// Split the input argument into a slice of strings:
func SplitInput(input string)[]string {
	return strings.Split(input, " ")
}