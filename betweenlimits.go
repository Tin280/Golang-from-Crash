package sprint 
func BetweenLimits(from, to rune) string {
	result := ""
	for p := from +1; p < to; p++ {
		result += string(p)
		}
	return result
}