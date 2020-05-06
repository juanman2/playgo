// Copyright 2020 Juan Tellez All rights reserved.

package strutil

import (
	"strings"
)

// WordCapitalize changes a word so the first letter is always capitalized
func WordCapitalize(word string) string {
	var n []string
	for i, r := range word {
		if i == 0 {
			n = append(n, strings.ToUpper(string(r)))
		} else {
			n = append(n, strings.ToLower(string(r)))
		}
	}
	return strings.Join(n, "")
}

// SentenceToCamelCap changes a sentence into a single camel cap identifier.
func SentenceToCamelCap(sentence string) string {

	wordArray := strings.Fields(sentence)
	for i, w := range wordArray {
		if i == 0 {
			wordArray[i] = strings.ToLower(w)
		} else {
			wordArray[i] = WordCapitalize(w)
		}
	}
	return strings.Join(wordArray, "")
}
