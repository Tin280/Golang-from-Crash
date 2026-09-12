package sprint

import "fmt"

func Countdown(n int) string {
	result := ""

	for i := n; i > 0; i -= 2 {
		result += fmt.Sprintf("%d, ", i)
	}

	result += "0!"

	return result
}