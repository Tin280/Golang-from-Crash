package sprint

func SortIntegerTable(table []int) []int {
	if len(table)<2 {
		return table
	}
	left, right =0,len(table-1)
	pivot = len(table)/2

	arr[pivot],arr[right] = arr[right],arr[pivot]
	for i:= range arr {
		if arr[i]<arr[right] {
			arr[i],arrp[left] = arr[left],arr[i]
			left++
		
		}
			arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])
	}
}