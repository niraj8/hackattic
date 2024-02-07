package main

import (
	"bufio"
	"fmt"
	"niraj8/hackattic/katas/solution"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// kata/its_almost_compression
	// for scanner.Scan() {
	// 	input := scanner.Text()
	// 	result := solution.ItsAlmostCompression(input)
	// 	fmt.Println(result)
	// }

	// kata/sorting_json_lines
	// var lines []string
	// for scanner.Scan() {
	// 	input := scanner.Text()
	// 	lines = append(lines, input)
	// }
	// result := solution.SortingJsonLines(lines)
	// for _, line := range result {
	// 	fmt.Println(line)
	// }

	// kata/the_sum_of_things
	for scanner.Scan() {
		input := scanner.Text()
		result := solution.TheSumOfThings(input)
		fmt.Println(result)
	}

}
