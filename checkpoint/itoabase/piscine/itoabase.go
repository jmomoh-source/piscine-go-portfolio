package piscine

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}
	if value == 0 {
		return "0"
	}

	baseStr := "0123456789ABCDEF"
	isNegative := value < 0

	var res []byte
	n := value
	for n != 0 {
		rem := n % base
		if rem < 0 {
			rem = -rem
		}
		res = append(res, baseStr[rem])
		n /= base
	}

	if isNegative {
		res = append(res, '-')
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}
