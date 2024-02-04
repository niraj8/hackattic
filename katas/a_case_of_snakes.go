package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {

	pattern := regexp.MustCompile(`/[A-Z]+/`)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println(pattern.MatchString(input))

	}
}
