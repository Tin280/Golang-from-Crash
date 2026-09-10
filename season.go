package sprint
func Season(month string) string {
	winter := []string{"dec", "jan", "feb"}
	spring := []string{"mar", "apr", "may"}
	summer := []string{"jun","jul","aug"}
	autumn := []string{"sep","oct","nov"}
	switch month {
	case winter[0], winter[1], winter[2]:
		return "winter"
	case spring[0], spring[1], spring[2]:
		return "spring"
	case summer[0], summer[1], summer[2]:
		return "summer"
	case autumn[0], autumn[1], autumn[2]:
		return "autumn"
	default:
		return "invalid input: " + month
	}
}