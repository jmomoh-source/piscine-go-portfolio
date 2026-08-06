package piscine

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	isNegative := false
	if n < 0 {
		isNegative = true
	}

	var digits []byte
	num := n
	for num != 0 {
		digit := num % 10
		if digit < 0 {
			digit = -digit
		}
		digits = append(digits, byte('0'+digit))
		num /= 10
	}

	if isNegative {
		digits = append(digits, '-')
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return string(digits)
}
