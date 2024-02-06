package solution

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func YesItFizz() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		splits := strings.Split(input, " ")
		start, end := MustParseInt(splits[0]), MustParseInt(splits[1])
		for i := start; i <= end; i++ {
			switch {
			case i%15 == 0:
				fmt.Println("FizzBuzz")
			case i%3 == 0:
				fmt.Println("Fizz")
			case i%5 == 0:
				fmt.Println("Buzz")
			default:
				fmt.Println(i)
			}
		}
	}
}

func MustParseInt(s string) int {
	parsedInt, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return parsedInt
}
