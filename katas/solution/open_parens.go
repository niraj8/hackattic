package solution

func OpenParens(input string) string {
	openParenCount := 0

	for _, chr := range input {
		switch chr {
		case '(':
			openParenCount += 1
		case ')':
			openParenCount -= 1
		default:
		}
	}
	if openParenCount != 0 {
		return "no"
	}
	return "yes"
}
