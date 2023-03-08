package levenshtein

import (
	"fmt"
)

var memo map[string]int

func Lev(source, dest string) int {
	return lev([]rune(source), []rune(dest))
}

func lev(source, dest []rune) int {
	m := len(source)
	n := len(dest)

	// Create array
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

	fmt.Println(d)
	return d[m][n]
}

// LevRecursive : Returns the Levenshtein distance between two strings.
func LevRecursive(source, dest string) int {
	// Create the memoization map
	memo = make(map[string]int)
	return levRecursive([]rune(source), []rune(dest))
}

func levRecursive(source, dest []rune) int {
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
		l := levRecursive(source[1:], dest[1:])
		memo[s1+d1] = l
		return l
	}

	l1 := levRecursive(source[1:], dest)     // Remove first letter in source
	l2 := levRecursive(source, dest[1:])     // Add one letter to source
	l3 := levRecursive(source[1:], dest[1:]) // Replace first letter in source

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
