package main

import "fmt"

func appendReturnsANewSlice() {
	// britney.pets = animals
	animals := []string{"cat", "dog"}

	type person struct {
		pets []string
	}
	var britney person
	animals = append(animals, "bear", "owl")
	britney.pets = animals

	animals[0] = "CHARMANDER"
	fmt.Println(animals, britney.pets)

	// britney.pets = append(animals, "bear", "owl")
}
