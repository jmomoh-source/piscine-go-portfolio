package piscine

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	m1 := make(map[rune]bool)
	for _, r := range str1 {
		m1[r] = true
	}

	m2 := make(map[rune]bool)
	for _, r := range str2 {
		m2[r] = true
	}

	count := 0
	for r := range m1 {
		if !m2[r] {
			count++
		}
	}
	for r := range m2 {
		if !m1[r] {
			count++
		}
	}

	return count
}
