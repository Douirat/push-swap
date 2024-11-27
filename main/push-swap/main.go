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
	fmt.Println(args)
	if err := outils.SplitInput(stacks, args[0]); err != nil {
		fmt.Println("Error")
		return
	}
	stacks.Sort()
	stacks.DisplayStackA()
	stacks.DisplayStackB()
	stacks.DisplayOperations()
	
}
