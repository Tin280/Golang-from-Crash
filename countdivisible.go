package sprint
func CountDivisible(from, to, step, divisor int) int {
	result := 0
	switch {
	case step <=0 || divisor ==0:
		return 0
	default:
		for i :=from;i<to; i += step{
			if( i%divisor==0){
				result++
			}
		}
	}
	return result
}