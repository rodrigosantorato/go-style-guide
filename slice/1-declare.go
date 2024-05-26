package main

import "fmt"

func declareEmptySlice() {
	makeS := make([]int, 0)
	inferS := []int{}
	var nilS []int

	fmt.Println(makeS, inferS, nilS)
	fmt.Println(nilS == nil)

	// ideal
}

func nilIsAValidSlice() []int {
	return nil
}
