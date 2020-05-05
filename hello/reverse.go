/*
Utilities to manipulate strings.
*/
package main

import (
	"fmt"
	"strings"
)

// Reverse changes the order of a string
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// SayHello prints attributes bout the parameter passed in
func SayHello(foo string) {
	fmt.Printf("%s contains ll %t\n", foo, strings.Contains(foo, "hello"))
	fmt.Printf("reverse %s\n", Reverse(foo))
	fmt.Printf("%s contains ll %t\n", foo, strings.Contains(foo, "hello"))
}


