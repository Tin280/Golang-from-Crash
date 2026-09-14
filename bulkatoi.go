package main 
// import "fmt"

// func StrToInt(s string) int {
// 	if len(s) == 0 {
// 		return 0
// 	}

// 	sign := 1
// 	start := 0

// 	if s[0] == '+' {
// 		start = 1
// 	} else if s[0] == '-' {
// 		sign = -1
// 		start = 1
// 	}

// 	// Chỉ có dấu + hoặc -
// 	if start == len(s) {
// 		return 0
// 	}

// 	result := 0

// 	for i := start; i < len(s); i++ {
// 		if s[i] < '0' || s[i] > '9' {
// 			return 0
// 		}

// 		result = result*10 + int(s[i]-'0')
// 	}

// 	return result * sign
// }
func BulkAtoi(arr []string) any {
	var result [] int
	for _, r := range arr {
			result = append(result,StrToInt(r))

	}
	return result
}

// func main(){
// 	fmt.Println(BulkAtoi([]string{"8", "kood", "-13"}))
// }