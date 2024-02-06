package main

import (
	"bufio"
	"fmt"
	"niraj8/hackattic/katas/solution"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		result := solution.ItsAlmostCompression(input)
		fmt.Println(result)
	}

}
