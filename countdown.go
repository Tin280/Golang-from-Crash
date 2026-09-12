package sprint

import "strconv"

func Countdown(n int) string {
	result := ""

	for i := n; i > 0; i -= 2 {
		result += strconv.Itoa(i) + ", "
	}

	result += "0!"

	return result
}