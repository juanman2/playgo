// Copyright 2020 Juan Tellez All rights reserved.

package strutil

import "strings"

// IsPalindrome returns true if the string is a valid palindrome
func IsPalindrome(str string) bool {
	r := []rune(str)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}

// CouldBePalindrome returns true if the string is a valid palindrome, when
// you remove the rune indexed by idx
func CouldBePalindrome(str string) bool {
	r := []rune(str)

	f := func(str string, lower bool) bool {
		jInc, iInc := 0, 0
		for i, j := 0, len(r)-1; i < (len(r) - 1)/2; i, j = i+1, j-1 {

			// skip this i index on the way up?
			if (lower == true && iInc == 0 && r[i] != r[j]) {
				iInc = 1
			}

			// skip this j index on the way down?
			if (lower == false && jInc == 0 && r[i] != r[j]) {
				jInc = -1
			}
			
			if r[i + iInc] != r[j + jInc] {
				return false
			}
		}
		return true
	}

	if f(str, true) == true || f(str,false) == true {
		return true
	}

	return false
}

// ValidPalindrome makes a palindrome if possible out of the string passed in
// by removing a single character from the string, otherwise returns the empty
// string
func validPalindrome(str string) bool {
	if CouldBePalindrome(str) == true {
		return true
	}
	return false
}

// ValidPalindrome makes a palindrome if possible out of the string passed in
// by removing a single character from the string, otherwise returns the empty
// string
func ValidPalindrome(str string) string {
	for i := range str  {
		p := TrimIth(str, i)
		if IsPalindrome(p) == true {
			return p
		}
	}
	return ""
}

// TrimIth removes the ith rune from the string
func TrimIth(str string, idx int) string {
	
	var newStr []string
	for i, r := range str {
		if idx != i {
			newStr = append(newStr, string(r))
		}
	}
	return strings.Join(newStr, "")
}


