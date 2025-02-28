package main

import "fmt"

func main() {
	rows := 7
	cols := 8

	fmt.Println("Enter the number of rows:")
	fmt.Scan(&rows)
	fmt.Println("Enter the number of seats in each row:")
	fmt.Scan(&cols)

	seats := rows * cols
	income := 0

	if seats <= 60 {
		income = seats * 10
	} else {
		firstHalf := rows / 2
		secondHalf := rows - firstHalf
		income = cols * (firstHalf*10 + secondHalf*8)
	}
	fmt.Println("Total income:\n")
	fmt.Printf("$%d", income)
}
