package piscine

func CanJump(steps []uint) bool {
	if len(steps) == 0 {
		return false
	}
	if len(steps) == 1 {
		return true
	}
	idx := 0
	last := len(steps) - 1
	for idx < last {
		s := int(steps[idx])
		if s == 0 {
			return false
		}
		idx += s
		if idx > last {
			return false
		}
		if idx == last {
			return true
		}
	}
	return idx == last
}
