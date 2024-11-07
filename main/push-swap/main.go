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
	// stackB := s.NewStack()
	args := os.Args[1:]
	// // Create a scanner to read from stdin
	// scanner := bufio.NewScanner(os.Stdin)
	// // Continuously read lines from stdin
	// for scanner.Scan() {
	// 	// Pri SplitInput(stack *p.Stack, input string)rgs) == 0 {
	// 	fmt.Println("Not enough arguments...")
	// 	return
	// }
	outils.SplitInput(stacks, args[0])
	fmt.Println(stacks.StackA)
}
