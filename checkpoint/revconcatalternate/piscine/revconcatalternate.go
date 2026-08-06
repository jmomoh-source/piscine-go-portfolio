package piscine

func RevConcatAlternate(slice1, slice2 []int) []int {
	s1 := make([]int, len(slice1))
	for i, v := range slice1 {
		s1[len(slice1)-1-i] = v
	}
	s2 := make([]int, len(slice2))
	for i, v := range slice2 {
		s2[len(slice2)-1-i] = v
	}

	var res []int
	i, j := 0, 0
	if len(s1) > len(s2) {
		diff := len(s1) - len(s2)
		res = append(res, s1[:diff]...)
		i = diff
	} else if len(s2) > len(s1) {
		diff := len(s2) - len(s1)
		res = append(res, s2[:diff]...)
		j = diff
	}

	for i < len(s1) || j < len(s2) {
		if i < len(s1) {
			res = append(res, s1[i])
			i++
		}
		if j < len(s2) {
			res = append(res, s2[j])
			j++
		}
	}

	return res
}
