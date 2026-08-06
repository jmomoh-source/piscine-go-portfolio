package piscine

import "fmt"

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}

	var res string
	i := 0
	for i < len(s) {
		ch := s[i]
		count := 0
		j := i
		for j < len(s) && s[j] == ch {
			count++
			j++
		}
		res += fmt.Sprintf("%d%c", count, ch)
		i = j
	}
	return res
}
