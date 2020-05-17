package playgo

import "fmt"

/*
Given an input string (s) and a pattern (p), implement regular expression
matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z,
and characters like . or *.

NOTE NOTE:
Using a DP approach lifted from Tushar Roy:
https://www.youtube.com/channel/UCZLJf_R2sWyUtXSKiKlyvAw
*/

func isMatch(s string, p string) bool {
	return IsRegexpMatchDP(s, p)
}

// IsRegexpMatchDP uses dynamic programming technique to solve the
// regexp puzzle
func IsRegexpMatchDP(s string, p string) bool {

	T := make([][]bool, len(s)+1)
	for i := range T {
		T[i] = make([]bool, len(p)+1)
	}

	T[0][0] = true

	for i := 1; i < len(p)+1; i = i + 1 {
		if i > 1 && p[i-1] == '*' {
			T[0][i] = T[0][i-2]
		}
	}

	for i := 1; i < len(s)+1; i = i + 1 {
		for j := 1; j < len(p)+1; j = j + 1 {
			fmt.Printf("(%d,%d) .. ", i, j)
			switch {
			case p[j-1] == '.' ||
				p[j-1] == s[i-1]:
				T[i][j] = T[i-1][j-1]
			case p[j-1] == '*':
				T[i][j] = T[i][j-2]
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					T[i][j] = T[i][j] || T[i-1][j]
				}
			default:
				T[i][j] = false
			}
		}
	}

	return T[len(s)][len(p)]
}
