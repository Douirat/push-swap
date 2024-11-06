package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// outils "github.com/Douirat/push-swap/logic/outils"
)

func main() {
	args := os.Args
	// Create a scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)
	flags := 
	// Continuously read lines from stdin
	for scanner.Scan() {
		// Print the received line from stdin
		fmt.Println("Received input:", scanner.Text())

	}
	if len(args) == 0 {
		fmt.Println("Not enough arguments...")
		return
	}
	// slc := outils.SplitInput(args[0])
	fmt.Println(args)
}
