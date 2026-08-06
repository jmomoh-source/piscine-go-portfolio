package piscine

func ConcatAlternate(slice1, slice2 []int) []int {
	first, second := slice1, slice2
	if len(slice2) > len(slice1) {
		first, second = slice2, slice1
	}

	var res []int
	i, j := 0, 0
	for i < len(first) || j < len(second) {
		if i < len(first) {
			res = append(res, first[i])
			i++
		}
		if j < len(second) {
			res = append(res, second[j])
			j++
		}
	}
	return res
}
