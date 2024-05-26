package main

import "fmt"

func aSliceIsAReferenceToAnUnderlyingArray() {
	// pode dar bug
	bands := []string{"muse", "bad omens", "architects"}

	type roquista struct {
		bands []string
	}
	var brother roquista
	brother.bands = bands

	bands[0] = "AVIÕES DO FORRÓ"
	fmt.Println(bands, brother.bands)

	// ideal
}
