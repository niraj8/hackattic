package solution

import (
	"log"
	"strconv"
	"strings"
)

func TheSumOfThings(input string) int64 {
	parts := strings.Split(input, " ")
	var sum int64 = 0
	for _, part := range parts {
		var (
			num int64
			err error
		)
		switch {
		// hex, octal, binary
		case strings.HasPrefix(part, "0x") || strings.HasPrefix(part, "0o") || strings.HasPrefix(part, "0b"):
			num, err = strconv.ParseInt(part, 0, 64)
			if err != nil {
				log.Println(err)
			}
		// decimal & characters
		default:
			i, err := strconv.Atoi(part)
			if err != nil {
				// log.Println("error converting %w to int: %w", part, err)
				i = int(part[0])
			}
			num = int64(i)
		}
		sum += num
	}

	return sum
}
