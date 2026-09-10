package sprint
func IntVsFloat(i int, f float32) string {
	switch {	
	case float32(i) > (f):
		return "Integer"
	case float32(i) < (f):
		return "Float"
	default:
		return "Same"
	}

}