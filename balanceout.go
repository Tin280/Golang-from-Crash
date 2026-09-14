package sprint
func BalanceOut(arr []bool) []bool {
	trueCount := 0
	falseCount := 0

	for _, value := range arr {
		if value {
			trueCount++
		} else {
			falseCount++
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