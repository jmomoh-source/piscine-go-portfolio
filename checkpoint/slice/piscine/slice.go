package piscine

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 {
		return nil
	}

	start := nbrs[0]
	end := len(a)
	if len(nbrs) > 1 {
		end = nbrs[1]
	}

	if start < 0 {
		start = len(a) + start
	}
	if end < 0 {
		end = len(a) + end
	}

	if start < 0 {
		start = 0
	}
	if end > len(a) {
		end = len(a)
	}

	if start >= end || start >= len(a) || end < 0 {
		return nil
	}

	return a[start:end]
}
