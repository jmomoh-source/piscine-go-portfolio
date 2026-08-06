package piscine

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return s
	}

	for i := 0; i < len(s); i++ {
		isUpper := s[i] >= 'A' && s[i] <= 'Z'
		isLower := s[i] >= 'a' && s[i] <= 'z'

		if !isUpper && !isLower {
			return s
		}

		if i > 0 && isUpper && (s[i-1] >= 'A' && s[i-1] <= 'Z') {
			return s
		}
		if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
			return s
		}
	}
	
	var result []rune 
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return string(result)
}