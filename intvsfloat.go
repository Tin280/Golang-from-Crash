package sprint
func IntVsFloat(i int, f float32) string {
	switch {	
	case i > int(f):
		return "Interger"
	case i < int(f):
		return "Float"
	default:
		return "Same"
	}

}