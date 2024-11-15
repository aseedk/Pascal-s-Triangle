# Pascal's Triangle Generator

This Go program generates Pascal's Triangle up to a specified number of rows and prints it to the console. Pascal's Triangle is a triangular array of numbers where each number is the sum of the two numbers directly above it.

## Purpose

The purpose of this program is to generate and display Pascal's Triangle, which has applications in combinatorics, algebra, and probability theory. The program demonstrates basic iterative logic in Go and provides a simple example of using arrays and loops to build mathematical structures.

## How the Program Works

- The program iterates from row 0 to row `n` (where `n` is predefined as `10` in the code).
- The first two rows are initialized directly. From the third row onwards, each new row is constructed using the values of the previous row.
- Each row begins and ends with the number `1`, and every other value is the sum of two values from the previous row.
- The rows are printed to the console, and the final row is highlighted separately.

## Expected Input

- The number of rows `n` to generate is defined within the code. 
- Currently, `n` is set to `10`. To change the number of rows, modify the value of `n` in the code.

## Expected Output

The program prints Pascal's Triangle row by row. For `n = 10`, the output is:


Note: The first row is printed as an empty slice `[]` for readability.

## How to Run

1. Make sure you have [Go installed](https://golang.org/doc/install) on your system.
2. Save the program in a file named `main.go`.
3. Open a terminal and navigate to the directory containing `main.go`.
4. Run the program using the command:

   ```bash
   go run main.go

This `README.md` file provides an overview of the program's purpose, functionality, expected input and output, and instructions for running it, along with suggestions for future improvements.
