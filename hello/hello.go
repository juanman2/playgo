/*
Testing module, new lsp emacs setup, and refreshing Go
*/
package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func hello() {
	foo := string("hello")
	fmt.Printf("%s contains ll %t\n", foo, strings.Contains(foo, "hello"))
	fmt.Printf("reverse %s\n", reverse(foo))
	fmt.Printf("%s contains ll %t\n", foo, strings.Contains(foo, "hello"))
}

func main() {
	hello()
}
