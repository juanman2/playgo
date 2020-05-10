// Copyright 2020 Juan Tellez All rights reserved.

package strutil

import "strings"

// IsPalindrome returns true if the string is a valid palindrome
func IsPalindrome(str string) bool {
	rev := Reverse(str)
	if (str == rev) {
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
