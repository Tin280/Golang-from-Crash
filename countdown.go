package sprint

func Countdown(n int) string {
	result := ""

	for i := n; i > 0; i -= 2 {
		result += string(rune(i+'0')) + ", "
	}

	result += "0!"

	return result
}