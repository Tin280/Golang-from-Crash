package sprint
func Season(month string) string {
	winter := []string{"dec", "jan", "feb"}
	spring := []string{"mar", "apr", "may"}
	summer := []string{"jun","jul","aug"}
	autumn := []string{"sep","oct","nov"}
	switch month {
	case winter[0], winter[1], winter[2]:
		return "Winter"
	case spring[0], spring[1], spring[2]:
		return "Spring"
	case summer[0], summer[1], summer[2]:
		return "Summer"
	case autumn[0], autumn[1], autumn[2]:
		return "Autumn"
	default:
		return "Invalid input: " + month
	}
}