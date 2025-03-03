package main

import "fmt"

func printSeating(cinema [][]rune) {
	fmt.Println("Cinema:")
	fmt.Print(" ")
	for i := 1; i <= len(cinema[0]); i++ {
		fmt.Printf(" %d", i)
	}
	fmt.Println()
	for i, row := range cinema {
		fmt.Printf("%d", i+1)
		for _, seat := range row {
			fmt.Printf(" %c", seat)
		}
		fmt.Println()
	}
}

func main() {
	var rows, seats int

	fmt.Println("Enter the number of rows:")
	fmt.Scan(&rows)
	fmt.Println("Enter the number of seats in each row:")
	fmt.Scan(&seats)

	cinema := make([][]rune, rows)
	for i := range cinema {
		cinema[i] = make([]rune, seats)
		for j := range cinema[i] {
			cinema[i][j] = 'S'
		}
	}

	for {
		fmt.Println("1. Show the seats")
		fmt.Println("2. Buy a ticket")
		fmt.Println("0. Exit")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			printSeating(cinema)
		case 2:
			var rowNumber, seatNumber int
			fmt.Println("Enter a row number:")
			fmt.Scan(&rowNumber)
			fmt.Println("Enter a seat number in that row:")
			fmt.Scan(&seatNumber)

			if cinema[rowNumber-1][seatNumber-1] == 'B' {
				fmt.Println("The ticket has already been purchased!")
				continue
			}

			totalSeats := rows * seats
			var ticketPrice int

			if totalSeats <= 60 {
				ticketPrice = 10
			} else {
				firstHalf := rows / 2
				if rowNumber <= firstHalf {
					ticketPrice = 10
				} else {
					ticketPrice = 8
				}
			}
			fmt.Printf("Ticket price: $%d\n", ticketPrice)

			cinema[rowNumber-1][seatNumber-1] = 'B'
		case 0:
			return
		default:
			fmt.Println("Invalid option, please choose again.")
		}
	}
}
