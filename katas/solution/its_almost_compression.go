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
		} else {
			compressed += processCharacter(prev, counter)
			prev = chr
			counter = 1
		}
	}
	compressed += processCharacter(prev, counter)
	return compressed
}

func processCharacter(c rune, counter int) string {
	switch counter {
	case 1:
		return string(c)
	case 2:
		return string(c) + string(c)
	default:
		return fmt.Sprint(counter) + string(c)
	}
}
