// Average2 calculates the average of several numbers.

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Use slice to omit program name
	// Assign slice to arguments variable
	arguments := os.Args[1:]

	// Declare sum variable float64 to hold sum of arguments
	var sum float64 = 0

	// Use for range loop to iterate over slice and convert to float64 and add to sum
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}

	// Declare sampleCount variable to hold number of arguments
	sampleCount := float64(len(arguments))

	// Declare and assign average variable to hold sum / sampleCount

	fmt.Printf("Average: %0.2f\n", sum/sampleCount)

}
