package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Douirat/push-swap/logic/checker"
	"github.com/Douirat/push-swap/logic/push-swap"
	"github.com/Douirat/push-swap/logic/utils"
)

func main() {
	stacks := pushswap.Init()
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println()
		return
	}
	if err := utils.SplitInput(stacks, args[0]); err != nil {
		fmt.Println("Error")
		return
	}
	// Reading stdin
	// Check if stdin is empty
	stat, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("Error checking stdin:", err)
	} else {
		// If stdin is not being piped or redirected, handle it
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			fmt.Println("Ko")
			return
		}

		scanner := bufio.NewScanner(os.Stdin)
		Instructions := []string{}
		fmt.Println("START")
		for scanner.Scan() {
			if scanner.Text() != "" {
				Instructions = append(Instructions, scanner.Text())
			}
		}
		fmt.Println("END")
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading stdin:", err)
		}
		if len(Instructions) == 0 {
			fmt.Println("Ko")
			return
		} else {
			stacks.Operations = Instructions
		}
		fmt.Printf("the stack is: %s\n", stacks.Operations)
		isSorted := checker.Check(stacks, Instructions)
		if isSorted {
			fmt.Println("OK")
		} else {
			fmt.Println("KO")
		}
	}
}
