package main

import (
	"fmt"
	"github.com/Douirat/push-swap/logic/utils"
	"github.com/Douirat/push-swap/logic/push-swap"
	"os"
)

func main() {
	stacks := pushswap.Init()
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println()
		return
	}
	if err := utils.SplitInput(stacks, args[0]); err != nil {
		fmt.Println("Error")
		return
	}
	stacks.Sort()
	stacks.DisplayOperations()
}
