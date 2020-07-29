package playgo

import "fmt"

/*
Given an input string (s) and a pattern (p), implement regular expression
matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
'+' Matches one or more of the preceding element.
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z,
and characters like . or *.

NOTE NOTE:
Using a DP approach lifted from Tushar Roy:
https://www.youtube.com/channel/UCZLJf_R2sWyUtXSKiKlyvAw
*/

// IsRegexpMatchDPPlus uses dynamic programming technique to solve the
// regexp puzzle, including a +
func IsRegexpMatchDPPlus(s string, p string) bool {

	fmt.Printf("Testing DPPlus string: %s pattern: %s\n", s, p)

	T := make([][]bool, len(s)+1)
	for i := range T {
		T[i] = make([]bool, len(p)+1)
	}

	T[0][0] = true

	for j := 1; j < len(p); j = j + 1 {
		if p[j] == '*' {
			T[0][j+1] = T[0][j-1]
		}
	}
	fmt.Printf("%v\n", T[0])

	for i := 0; i < len(s); i = i + 1 {
		for j := 0; j < len(p); j = j + 1 {
			switch {
			case p[j] == '.' ||
				p[j] == s[i]:
				T[i+1][j+1] = T[i][j] || T[i+1][j]
			case p[j] == '*':
				T[i+1][j+1] = T[i+1][j-1]
				if p[j-1] == '.' || p[j-1] == s[i] {
					T[i+1][j+1] = T[i+1][j+1] || T[i][j+1]
				}
			case p[j] == '+':
				if p[j-1] == s[i] {
					T[i+1][j+1] = T[i][j] || T[i][j+1]
				} else {
					T[i+1][j+1] = T[i][j]
				}
			default:
				T[i+1][j+1] = false
			}
		}
		fmt.Printf("%v\n", T[i+1])
	}

	return T[len(s)][len(p)]
}
