package pushswap

// Create a structure t hold both of my structs:
type Stacks struct {
	StackA []int
	StackB []int
	TopA   int
	TopB   int
}

// Instantiate the stacks struct:
func Init() *Stacks {
	stacks := &Stacks{}
	stacks.TopA = -1
	stacks.TopB = -1
	return stacks
}

// 
