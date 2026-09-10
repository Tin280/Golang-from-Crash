package sprint
func Doop(a int, op string, b int) int {
	result := 0
	switch op {
	case "+":
		result = a+b
	case "-":
		result = a-b 
	case "*":
		result = a*b
	case "/":
		if b == 0 {
			result = 0
		} else {
			result = a/b
		}
	default: 
		result = 0
	}
		return result
}