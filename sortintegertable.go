package sprint

func SortIntegerTable(table []int) []int {
	if len(table)<2 {
		return table
	}
	left, right :=0,len(table-1)
	pivot := len(table)/2

	table[pivot],table[right] = table[right],table[pivot]
	for i:= range table {
		if table[i]<table[right] {
			table[i],table[left] = table[left],table[i]
			left++
		
		}
			table[left], table[right] = table[right], table[left]

	SortIntegerTable(table[:left])
	SortIntegerTable(table[left+1:])
	}
	return arr
}