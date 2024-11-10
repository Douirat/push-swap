package main

import (
	// "bufio"
	"fmt"
	"os"
	outils "github.com/Douirat/push-swap/logic/outils"
	s "github.com/Douirat/push-swap/logic/push-swap"
)

func main() {
	stacks := s.Init()
	args := os.Args[1:]
	outils.SplitInput(stacks, args[0])
	fmt.Println(stacks.StackA)
	fmt.Println(stacks.StackB)
	// stacks.StackB = []int{4, 2, 3, 5, 8}
	stacks.Sort()
	fmt.Println(stacks.StackA)
	fmt.Println(stacks.StackB)
}
