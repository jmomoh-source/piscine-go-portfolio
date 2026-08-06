package piscine

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	hasDot := false
	dotIndex := -1
	start := 0
	if dec[0] == '+' || dec[0] == '-' {
		start = 1
	}

	if start >= len(dec) {
		return dec + "\n"
	}

	digitCount := 0
	for i := start; i < len(dec); i++ {
		if dec[i] == '.' {
			if hasDot {
				return dec + "\n"
			}
			hasDot = true
			dotIndex = i
		} else if dec[i] >= '0' && dec[i] <= '9' {
			digitCount++
		} else {
			return dec + "\n"
		}
	}

	if digitCount == 0 || !hasDot {
		return dec + "\n"
	}

	afterDot := dec[dotIndex+1:]
	if len(afterDot) == 0 {
		return dec + "\n"
	}
	allZeroAfter := true
	for i := 0; i < len(afterDot); i++ {
		if afterDot[i] != '0' {
			allZeroAfter = false
			break
		}
	}
	if allZeroAfter {
		return dec + "\n"
	}

	raw := dec[:dotIndex] + dec[dotIndex+1:]

	sign := ""
	if raw[0] == '+' || raw[0] == '-' {
		sign = string(raw[0])
		if sign == "+" {
			sign = ""
		}
		raw = raw[1:]
	}

	k := 0
	for k < len(raw)-1 && raw[k] == '0' {
		k++
	}
	raw = raw[k:]

	return sign + raw + "\n"
}
