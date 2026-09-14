package main 
// import "fmt"
func BulkAtoi(arr []string) []int {
	var result []int

	for _, s := range arr {
		if len(s) == 0 {
			return nil
		}

		sign := 1
		start := 0

		if s[0] == '+' {
			start = 1
		} else if s[0] == '-' {
			sign = -1
			start = 1
		}

		if start == len(s) {
			return nil
		}

		num := 0

		for i := start; i < len(s); i++ {
			if s[i] < '0' || s[i] > '9' {
				return nil
			}

			num = num*10 + int(s[i]-'0')
		}

		result = append(result, num*sign)
	}

	return result
}
// func main(){
// 	fmt.Println(BulkAtoi([]string{"8", "kood", "-13"}))
// }