package solution

import "fmt"

func ItsAlmostCompression(input string) string {
	if len(input) == 0 {
		return ""
	}
	var (
		prev       rune   = rune(input[0])
		counter    int    = 1
		compressed string = ""
	)
	for _, chr := range input[1:] {
		if chr == prev {
			counter += 1
		} else if counter > 2 {
			compressed += fmt.Sprint(counter) + string(prev)
			prev = chr
			counter = 1
		} else if counter == 2 {
			compressed += string(prev) + string(prev)
			prev = chr
			counter = 1
		} else {
			compressed += string(prev)
			prev = chr
			counter = 1
		}
	}
	if counter > 2 {
		compressed += fmt.Sprint(counter) + string(prev)
	} else if counter == 2 {
		compressed += string(prev) + string(prev)
	} else {
		compressed += string(prev)
	}
	return compressed
}
