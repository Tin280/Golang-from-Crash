package sprint

func AlphaNumber(n int) string {
	negative := false

	if n < 0 {
		negative = true
		n = -n
	}

	result := ""

	if n == 0 {
		result = "a"
	}

	for n > 0 {
		digit := n % 10
		result = string(rune('a'+digit)) + result
		n /= 10
	}

	if negative {
		result = "-" + result
	}

	return result
}