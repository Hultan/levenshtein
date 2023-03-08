package levenshtein

// Lev : Returns the Levenshtein distance between two strings.
func Lev(source, dest string) int {
	return lev([]rune(source), []rune(dest))
}
func lev(source, dest []rune) int {
	// fmt.Println(source, dest)
	if len(source) == 0 {
		return len(dest)
	}
	if len(dest) == 0 {
		return len(source)
	}
	if source[0] == dest[0] {
		return Lev(string(source[1:]), string(dest[1:]))
	}

	return 1 + min(
		lev(source[1:], dest),     // Remove first letter in source
		lev(source, dest[1:]),     // Add one letter to source
		lev(source[1:], dest[1:]), // Replace first letter in source
	)
}

func min(lev1, lev2, lev3 int) int {
	// Assume the first value is the smallest
	m := lev1
	// Check the second value
	if lev2 < m {
		m = lev2
	}
	// Check the third value
	if lev3 < m {
		m = lev3
	}
	// Return the smallest
	return m
}
