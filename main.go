package main

import "fmt"

/*
Pascal's Triangle example:
1
1 1
1 2 1
1 3 3 1
1 4 6 4 1
1 5 10 10 5 1
*/

func main() {
	n := 10
	var previousRow []int
	var currentRow []int

	// Generate Pascal's Triangle up to 'n' rows
	for i := 0; i <= n; i++ {
		fmt.Println("Previous Row: ", previousRow)

		// Handle the first two rows separately
		if i == 0 {
			previousRow = []int{1}
		} else if i == 1 {
			previousRow = []int{1, 1}
		} else {
			// Generate the next row based on the previous row
			currentRow = []int{1}
			for j := 1; j < len(previousRow); j++ {
				currentRow = append(currentRow, previousRow[j-1]+previousRow[j])
			}
			currentRow = append(currentRow, 1)
			previousRow = currentRow
		}
	}

	// Print the final row
	fmt.Println("Final Row:", currentRow)
}
