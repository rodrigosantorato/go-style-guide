package main

import (
	"fmt"
	"strings"
)

func sliceLengthAndCapacity() {
	// ruim
	words := []string{"this", "is", "a", "very", "long", "sentence"}
	var res []string
	fmt.Println(res, len(res), cap(res))
	for i := range words {
		res = append(res, strings.ToUpper(words[i]))
		fmt.Println(res, len(res), cap(res))
	}

	// ideal
}
