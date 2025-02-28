package main

import "fmt"

func main() {
	rows := 7
	cols := 8

	fmt.Println("Cinema:")
	fmt.Print("  ")
	for c := 1; c <= cols; c++ {
		fmt.Printf("%d ", c)
	}
	fmt.Println()

	for r := 1; r <= rows; r++ {
		fmt.Printf("%d ", r)
		for c := 1; c <= cols; c++ {
			fmt.Print("S ")
		}
		fmt.Println()
	}

}
