package main

import (
	"fmt"
)

func main() {
	// 1. Option one to create a map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4347362",
		"White": "#000000",
	}
	// 2. Option two to create a map
	//var capitals map[string]string
	//// 3. Option three to create a map
	//teams := make(map[string]int)
	//
	//capitals["Ecuador"] = "Quito"
	//
	//teams["Colombia"] = 10

	printMap(colors)
	//fmt.Println(colors)
}

// Iterating over a map
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex for color: ", color, " is ", hex)
	}
}
