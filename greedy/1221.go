package greedy

func balancedStringSplit(s string) int {
	splits := 0
	lc := 0
	rc := 0
	for _, ch := range s {
		switch ch {
		case 'L':
			lc++
		case 'R':
			rc++
		}
		if lc == rc {
			splits++
			lc = 0
			rc = 0
		}
	}
	return splits
}
