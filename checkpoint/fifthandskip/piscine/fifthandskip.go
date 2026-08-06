package piscine

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}

	var clean []rune
	for _, r := range str {
		if r != ' ' {
			clean = append(clean, r)
		}
	}

	if len(clean) == 0 {
		return "\n"
	}
	if len(clean) < 5 {
		return "Invalid Input\n"
	}

	var res []rune
	count := 0
	for i := 0; i < len(clean); i++ {
		res = append(res, clean[i])
		count++
		if count == 5 {
			if i+1 < len(clean) {
				res = append(res, ' ')
				i++ // Skip 6th char
			}
			count = 0
		}
	}

	return string(res) + "\n"
}
