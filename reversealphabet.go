package sprint

func ReverseAlphabet(step int) string {
	if step <= 0 {
		step = 1
	}

	result := ""

	for c := 'z'; c >= 'a'; c -= rune(step) {
		result += string(c)
	}

	return result
}