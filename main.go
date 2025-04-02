package main

import (
	"bufio"
	"fmt"
	"os"
)

// Function to decode the sequence
func decodeSequence(s string) string {
	n := len(s) + 1
	result := make([]int, n)

	// Start with the smallest number (0 or 1)
	result[0] = 0

	// Assign values based on conditions
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			result[i+1] = result[i] - 1
		} else if s[i] == 'R' {
			result[i+1] = result[i] + 1
		} else { // '=' case
			result[i+1] = result[i]
		}
	}

	// Normalize to ensure all numbers are non-negative
	minVal := 0
	for _, num := range result {
		if num < minVal {
			minVal = num
		}
	}
	for i := range result {
		result[i] -= minVal // Shift to make all values positive
	}

	// Convert result array to a string
	output := ""
	for _, num := range result {
		output += fmt.Sprintf("%d", num)
	}

	return output
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the sequence (L/R/=): ")
	scanner.Scan()
	input := scanner.Text()

	output := decodeSequence(input)
	fmt.Println("Decoded Sequence:", output)
}
