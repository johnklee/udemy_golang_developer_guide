package main

import "fmt"

func ex1() {
	// way1 to create map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}
	fmt.Println(colors)
}

func ex2() {
	// way2 to create map
	colors := make(map[string]string)
	colors["red"] = "#ff0000"
	colors["green"] = "#00ff00"
	delete(colors, "red")
	fmt.Println(colors) // map[green:#00ff00]
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("color=%s with hex=%s\n", color, hex)
	}
}

func ex3() {
	// How to iterate items of map in for loop
	color := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	printMap(color)
}

func main() {
	ex3()
}
