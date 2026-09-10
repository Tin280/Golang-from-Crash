package sprint

func BetweenLimits(from, to rune) string {
    result := ""
	switch {
	case from < to:
    	for p := from +1; p < to; p++ {
        	result += string(p)
    	}
	default:
		for p := to +1; p < from; p++ {
			result += string(p)
		}
	}
    return result
}