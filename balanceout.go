package sprint
func BalanceOut(arr []bool) []bool {
	truecount :=0
	falsecount :=0
	for _, value:= range arr{
		if value {
			truecount++
		} else {
			falsecount++
		}
	}
	if trueCount < falseCount {
		for i := 0; i < falseCount-trueCount; i++ {
			arr = append(arr, true)
		}
	} else if falseCount < trueCount {
		for i := 0; i < trueCount-falseCount; i++ {
			arr = append(arr, false)
		}
	}
	return arr
}