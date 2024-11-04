package main

import(
"github.com/Douirat/push-swap/logic"
)

func main() {
	arr :=[]int{4, 3, 12, 1, 0, 0, 5, 5, 3, 9};
	logic.PegeonSort(arr, len(arr))
	
}