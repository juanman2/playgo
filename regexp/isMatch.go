package playgo

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
This program doesn't work for all solutions.  Basically there are
too many permutations of how x* or .* can be applied, and some will work
while other don't. So in theory this code would need to be run multiple
times to compute the posibilities.  Instead it is better
to use a Dinamic Programming approach.
*/

import "fmt"

const (
	beginState       = iota
	dotState         = iota
	starState        = iota
	simpleMatchState = iota
	endState         = iota
	errorState       = iota
)

// DotMatch advances the regexp if it matches any single character.  If s
// is too short, it returns error
func DotMatch(s string, i int) (int, int) {
	if len(s) > i {
		return beginState, i + 1
	}
	return errorState, i
}

// SimpleMatch must match the string p, a single time
// on the string s
func SimpleMatch(s string, p byte, sIdx int) (int, int) {
	if sIdx < len(s) {
		if s[sIdx] == p {
			return beginState, sIdx + 1
		}
		return errorState, sIdx
	}
	return endState, sIdx
}

// StarMatch must match the string p, 0 or more times on the
// string s, if s is too short, it still matches
func StarMatch(s string, p byte, sIdx int) (int, int) {

	if s == "" {
		return beginState, sIdx
	}

	for i := sIdx; i < len(s); i = i + 1 {
		if p != '.' && s[i] != p {
			return beginState, i
		}
	}
	return beginState, len(s)
}

// RegexpState Returns the current state for the Regexp matching engine
func RegexpState(p string, regIdx int) (int, int, byte) {

	if regIdx < len(p) {

		switch {
		case p[regIdx] == '.':
			if regIdx+1 < len(p) && p[regIdx+1] == '*' {
				return starState, regIdx + 2, p[regIdx]
			}
			return dotState, regIdx + 1, p[regIdx]
		case (p[regIdx] >= 'a' && p[regIdx] <= 'z'):
			if regIdx+1 < len(p) && p[regIdx+1] == '*' {
				return starState, regIdx + 2, p[regIdx]
			}
			return simpleMatchState, regIdx + 1, p[regIdx]
		case p[regIdx] == '*':
			if regIdx > 0 {
				return starState, regIdx + 1, p[regIdx-1]
			}
			return errorState, regIdx, 0
		default:
			println(fmt.Errorf("Unexpected regexp character: %v",
				p[regIdx]))
			return errorState, regIdx, 0
		}
	}
	return endState, regIdx, 0
}

// IsRegexpMatch returns the next state for the regexp engine
func IsRegexpMatch(s string, p string) bool {

	c := byte(0)
	for pIdx, sIdx, state := 0, 0, beginState; state != errorState; {
		switch state {
		case dotState:
			state, sIdx = DotMatch(s, sIdx)
		case starState:
			state, sIdx = StarMatch(s, c, sIdx)
		case simpleMatchState:
			state, sIdx = SimpleMatch(s, c, sIdx)
		case beginState:
			if sIdx >= len(s) {
				state = endState
			} else {
				state, pIdx, c = RegexpState(p, pIdx)
			}
		case endState:
			if sIdx < len(s) {
				fmt.Printf("Unprocessed string %s %d\n", s, sIdx)
				state = errorState
			} else if pIdx < len(p) {
				fmt.Printf("Unprocessed regexp %s\n", p)
				state = errorState
			} else {
				return true
			}
		case errorState:
			fmt.Println(
				fmt.Errorf(
					"error processing the Regexp %s at %d on string %s at %d",
					p, pIdx, s, sIdx))
			return false
		default:
			panic(fmt.Sprintf("Unknown State: %d", state))
		}
	}
	return false
}
