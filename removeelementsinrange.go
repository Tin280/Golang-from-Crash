

package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	if from > to {
		from, to = to, from
	}

	if from < 0 {
		from = 0
	}

	if to > len(arr) {
		to = len(arr)
	}

	result := append([]float64{}, arr[:from]...)
	result = append(result, arr[to:]...)

	return result
}

