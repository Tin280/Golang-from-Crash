package sprint

func TimeConverter(totalSeconds int) (int, int, int) {
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	return hours, minutes, seconds
}