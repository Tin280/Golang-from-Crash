package sprint

func ShiftBy(r rune, step int) rune {
	base := 'a'
	offset := ((int(r -base)+step)%26+26)%26
	return rune(base + offset)
}