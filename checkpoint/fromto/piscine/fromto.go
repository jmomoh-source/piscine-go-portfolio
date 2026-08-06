package piscine

import "fmt"

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	var result string
	
	if from <= to {
		for  i := from;  i <= to;  i++ {
			result += fmt.Sprintf("%02d", i)

			if i != to {
				result += ", "
			}
		}
	} else {
		for i:= from; i >= to; i-- {
			result += fmt.Sprintf("%02d", i)

			if i != to {
				result += ", "
			}
		}
	}
	result += "\n"
	return result
}