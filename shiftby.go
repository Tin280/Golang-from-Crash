package sprint

import "fmt"

func ShiftBy(r rune, step int) rune {
	base := 'a'
	offset := ((int(r -base)+step)%26+26)%26
	result := base + rune(offset)
	return result
}
func main() {
	fmt.Println(string(ShiftBy('a', 4))) // Output: e
	fmt.Println(string(ShiftBy('z', 1))) // Output: a
	fmt.Println(string(ShiftBy('x', 5))) // Output: c
}