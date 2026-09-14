package sprint
func GenerateRange(min, max int) []int {
	var result [] int
	if (min>=max) {
		return nil
	}
	for min< max {
		result  = append(result,min)
		min +=1
	}
	return result
}


