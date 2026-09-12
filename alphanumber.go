package sprint

func AlphaNumber(n int) string {
	result := ""

	if n < 0 {
		result += "-"
		n = -n
	}

	if n == 0 {
		return "a"
	}

	for n > 0 {
		digit := n % 10
		result = string(rune('a'+digit)) + result
		n /= 10
	}

	return result
}