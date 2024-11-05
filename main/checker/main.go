package main

import (
	"fmt"
	"os"

	// "github.com/Douirat/push-swap/logic"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)
	fmt.Printf("%T", args)
}