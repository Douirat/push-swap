package main

import (
	"fmt"
	"os"

	outils "github.com/Douirat/push-swap/logic/outils"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)
	slc := outils.SplitInput(args[0])
	if len(slc) == 0 {
		fmt.Println("Not enough arguments...")
	}
	fmt.Println(len(slc))
}