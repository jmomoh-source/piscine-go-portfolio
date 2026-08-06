package piscine

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}

	var res []rune
	runes := []rune(arg)
	save := true
	for i := 0; i < len(runes); i += num {
		end := i + num
		if end > len(runes) {
			end = len(runes)
		}
		if save {
			res = append(res, runes[i:end]...)
		}
		save = !save
	}

	return string(res)
}
