package solution

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func AlmostBinary() {

	// 0 == . | 1 == #
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		binary := ""
		for _, c := range input {
			switch c {
			case '.':
				binary += "0"
			case '#':
				binary += "1"
			default:
			}
		}
		decimal, err := strconv.ParseInt(binary, 2, 64)
		if err != nil {
			log.Fatalf("failed to convert to decimal: %v", err)
		}
		fmt.Println(decimal)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
