package logic

import "fmt"

func PegeonSort(arr []int, n int) {
	fmt.Println(arr)
	var Max = arr[0]
	var Min = arr[0]
	for i:=0; i<len(arr); i++ {
		if arr[i] < Min {
			Min = arr[i]
		} else if arr[i] > Max {
			Max = arr[i]
		}
	}
	fmt.Printf("The max is: %d and the min %d\n", Max, Min)
	Range := (Max - Min) + 1
	fmt.Printf("The range is: %d\n", Range)
	pigeonHoles := make([]int, Range)
	for j:=0; j< n; j++ {
		pigeonHoles[(arr[j]-Min)]++
	}
	fmt.Println(pigeonHoles)
	for ind := 0; ind<Range; ind++ {
		for pigeonHoles[ind]-1 != 0 {
			arr[pigeonHoles[ind]+1] = ind + Min
		}
	}
	fmt.Println(arr)
 }