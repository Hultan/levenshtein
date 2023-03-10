package levenshtein

// Lev calculates the Levenshtein distance between two strings.
// See https://en.wikipedia.org/wiki/Levenshtein_distance for
// more information.
func Lev(source, dest string) int {
	return lev([]rune(source), []rune(dest))
}

// lev calculates the Levenshtein distance between two rune slices.
func lev(source, dest []rune) int {
	m := len(source)
	n := len(dest)

	d := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		d[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		d[i][0] = i
	}

	for j := 1; j <= n; j++ {
		d[0][j] = j
	}

	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			cost := 0
			if source[i-1] != dest[j-1] {
				cost = 1
			}

			d[i][j] = min(
				d[i-1][j]+1,      // deletion
				d[i][j-1]+1,      // insertion
				d[i-1][j-1]+cost, // substitution
			)
		}
	}

	return d[m][n]
}

// min returns the lowest value of the arguments.
// Example: min(3,5,1,7) returns 1.
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
