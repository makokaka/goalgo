package main

import (
	"log"
	"fmt"
)

func print_lcs(s string, r [][]int, i,j int) {
	if i == 0 || j == 0 {
		return
	}
	switch {
	case r[i][j] > r[i-1][j] && r[i][j] > r[i][j-1]:
		print_lcs(s, r, i-1, j-1)
		fmt.Printf("%c", s[i-1])
	case r[i][j-1] >= r[i-1][j]:
		print_lcs(s, r, i, j-1)
	default:
		print_lcs(s, r, i-1, j)
	}
}

func lcs(s1, s2 string) ([][]int) {
	m, n := len(s1), len(s2)
	var rec [][]int = make([][]int, m+1)
	for i, _ := range rec {
		rec[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		rec[i][0] = 0
	}
	for j := 0; j < n; j++ {
		rec[0][j] = 0
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			switch {
			case s1[i-1] == s2[j-1]:
				rec[i][j] = rec[i-1][j-1] + 1
			case rec[i][j-1] >= rec[i-1][j]:
				rec[i][j] = rec[i][j-1]
			default:
				rec[i][j] = rec[i-1][j]
			}
		}
	}
	log.Print(rec[m][n])
	print_lcs(s1, rec, m, n)
	return rec
}

func main() {
	lcs("abcd", "aebfcgd")
}
