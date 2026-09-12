package sprint

import "strconv"

func AlphaNumber(n int) string {
	num := strconv.Itoa(n)
	result := ""

	for i := 0; i < len(num); i++ {
		if num[i] == '-' {
			result += "-"
		} else {
			result += string('a' + (num[i] - '0'))
		}
	}

	return result
}