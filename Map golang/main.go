package main

import (
	"fmt"
)

func main() {
	//Make map 1
	// var colors map[string]string

	// Make map 2
	colors := make(map[string]string)
	colors["white"] = "sdfsdfg"
	colors["red"] = "sdfsdfg"
	colors["yellow"] = "sdfsdfg"

	// Make map 3
	// colors := map[string]string{
	// 	"black" : "#000000",
	// 	"red" : "#ff0000",
	// }

	//How to delete map
	// delete(colors, "white")

	//iterate map
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex of " + color + " is " + hex)
	}
}
