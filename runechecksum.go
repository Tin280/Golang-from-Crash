package sprint

func RuneChecksum(a, b rune) rune {
	posA := int(a-'a') + 1
	posB := int(b-'a') + 1

	result := (posA * posB) % 26

	if result == 0 {
		result = 26
	}

	return rune('a' + result - 1)
}