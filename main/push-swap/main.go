package main

import (
	"fmt"
	"github.com/Douirat/push-swap/logic/outils"
	"github.com/Douirat/push-swap/logic/push-swap"
	"os"
)

func main() {
	stacks := pushswap.Init()
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Not enough arguments!")
		return
	}
	outils.SplitInput(stacks, args[0])
	stacks.Sort()
	stacks.DisplayStackA()
	stacks.DisplayOperations()
}
