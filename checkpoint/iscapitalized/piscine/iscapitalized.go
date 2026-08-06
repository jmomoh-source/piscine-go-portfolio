package piscine

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}
	
	newWord := true

	for _, r := range s {
		if r == ' ' {
			newWord = true 
		} else if newWord {
			if r >= 'a' && r <= 'z' {
				return false
			}
			newWord = false
		}
	}
	return true
}
