package main

import (
	"os"
	outils "github.com/Douirat/push-swap/logic/outils"
	s "github.com/Douirat/push-swap/logic/push-swap"
)

func main() {
	stacks := s.Init()
	args := os.Args[1:]
	outils.SplitInput(stacks, args[0])
	stacks.Display()
}
