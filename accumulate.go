package sprint
func Accumulate(n int) int {
	switch {
	case n < 0:
		return 0
	default:
		sum := 0
		for i := 1; i <= n; i++ {
			sum += i
		}
		return sum
	}
}