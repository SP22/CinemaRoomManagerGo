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

func calculateTotalIncome(rows, seats int) int {
	totalSeats := rows * seats
	if totalSeats <= 60 {
		return totalSeats * 10
	}
	halfRows := rows / 2
	return (halfRows * seats * 10) + ((rows - halfRows) * seats * 8)
}

func calculateTicketPrice(rows int, seats int, rowNumber int) int {
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
	return ticketPrice
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

	purchasedTickets := 0
	currentIncome := 0
	totalIncome := calculateTotalIncome(rows, seats)

	for {
		fmt.Println("1. Show the seats")
		fmt.Println("2. Buy a ticket")
		fmt.Println("3. Statistics")
		fmt.Println("0. Exit")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			printSeating(cinema)
		case 2:
			for {
				var rowNumber, seatNumber int
				fmt.Println("Enter a row number:")
				fmt.Scan(&rowNumber)
				fmt.Println("Enter a seat number in that row:")
				fmt.Scan(&seatNumber)

				if rowNumber < 1 || rowNumber > rows || seatNumber < 1 || seatNumber > seats {
					fmt.Println("Wrong input!")
					continue
				}

				if cinema[rowNumber-1][seatNumber-1] == 'B' {
					fmt.Println("That ticket has already been purchased")
					continue
				}

				ticketPrice := calculateTicketPrice(rows, seats, rowNumber)
				fmt.Printf("Ticket price: $%d\n", ticketPrice)

				cinema[rowNumber-1][seatNumber-1] = 'B'
				purchasedTickets++
				currentIncome += ticketPrice
				break
			}
		case 3:
			percentage := float64(purchasedTickets) / float64(rows*seats) * 100
			fmt.Printf("Number of purchased tickets: %d\n", purchasedTickets)
			fmt.Printf("Percentage: %.2f%%\n", percentage)
			fmt.Printf("Current income: $%d\n", currentIncome)
			fmt.Printf("Total income: $%d\n", totalIncome)
		case 0:
			return
		default:
			fmt.Println("Invalid option, please choose again.")
		}
	}
}
