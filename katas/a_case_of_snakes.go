package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println(regex(input))
	}
	// isPrefixStripped := false
	// convertedString := ""
	// 	for _, c := range input {
	// 		chr := string(c)
	// 		if !isPrefixStripped {
	// 			_, err := strconv.Atoi(chr)
	// 			if strings.ToUpper(chr) == chr && err != nil {
	// 				isPrefixStripped = true
	// 				convertedString = convertedString + strings.ToLower(chr)
	// 			}
	// 		} else {
	// 			if strings.ToLower(chr) == chr {
	// 				convertedString = convertedString + chr
	// 			} else {
	// 				convertedString = convertedString + "_" + strings.ToLower(chr)
	// 			}
	// 		}
	// 	}
	// 	fmt.Println(convertedString)
	// }
}

func regex(input string) string {
	re := regexp.MustCompile("[a-zA-Z][a-z0-9]*")
	result := re.FindAllString(input, -1)
	if len(result[0]) <= 3 {
		result = result[1:]
	}
	return strings.ToLower(strings.Join(result, "_"))
}
