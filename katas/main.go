package main

import (
	"bufio"
	"fmt"
	"niraj8/hackattic/katas/solution"
	"os"
)

func main() {
	// solution.WhatDayWasIt()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		result := solution.ItsAlmostCompression(input)
		fmt.Println(result)
	}

}
