package levenshtein

var memo map[string]int

// Lev : Returns the Levenshtein distance between two strings.
func Lev(source, dest string) int {
	// Create the memoization map
	memo = make(map[string]int)
	return lev([]rune(source), []rune(dest))
}

func lev(source, dest []rune) int {
	s := string(source)
	d := string(dest)

	// Check memoization map
	m, ok := memo[s+d]
	if ok {
		return m
	}

	// Check empty source or dest
	if len(source) == 0 {
		memo[d] = len(dest)
		return memo[d]
	}
	if len(dest) == 0 {
		memo[s] = len(source)
		return memo[s]
	}

	s1 := string(source[1:])
	d1 := string(dest[1:])

	// Check if first char/rune is equal
	if source[0] == dest[0] {
		l := lev(source[1:], dest[1:])
		memo[s1+d1] = l
		return l
	}

	l1 := lev(source[1:], dest)     // Remove first letter in source
	l2 := lev(source, dest[1:])     // Add one letter to source
	l3 := lev(source[1:], dest[1:]) // Replace first letter in source

	memo[s1+d] = l1
	memo[s+d1] = l2
	memo[s1+d1] = l3

	return 1 + min(l1, l2, l3)
}

func min(lev ...int) int {
	// Assume the first value is the smallest
	m := lev[0]

	// Loop through the remaining values
	for i := 1; i < len(lev); i++ {
		if lev[i] < m {
			m = lev[i]
		}
	}

	// Return the smallest
	return m
}
