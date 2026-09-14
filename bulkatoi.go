package main

// import "fmt"

func BulkAtoi(arr []string) []int {
	var resultfinal []int

	for _, s := range arr {
		if len(s) == 0 {
			resultfinal = append(resultfinal, 0)
			continue
		}

		sign := 1
		start := 0

		if s[0] == '+' {
			start = 1
		} else if s[0] == '-' {
			sign = -1
			start = 1
		}

		// Chỉ có "+" hoặc "-"
		if start == len(s) {
			resultfinal = append(resultfinal, 0)
			continue
		}

		result := 0
		valid := true

		for i := start; i < len(s); i++ {
			if s[i] < '0' || s[i] > '9' {
				valid = false
				break
			}

			result = result*10 + int(s[i]-'0')
		}

		if !valid {
			resultfinal = append(resultfinal, 0)
			continue
		}

		resultfinal = append(resultfinal, result*sign)
	}

	return resultfinal
}

// func main() {
// 	fmt.Println(BulkAtoi([]string{"8", "kood", "-13"}))
// }