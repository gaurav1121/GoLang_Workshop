package main

import (
	"fmt"
	// s "github.com/inancgumus/prettyslice"
)

func main() {

	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}

	// s.MaxPerLine = 4
	// s.Show("All items", items)

	// top3 := items[:3]
	// s.Show("Top 3 items", top3)

	// l := len(items)

	// last4 := items[l-4:]
	// s.Show("Last 4 items", last4)

	// mid := last4[1:3]
	// s.Show("Last4[1:3]", mid)

	fmt.Printf("slicing : %T %[1]q\n", items[2:3])
	fmt.Printf("indexing: %T %[1]q\n", items[2])
}
