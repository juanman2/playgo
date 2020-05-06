/*
Testing module, new lsp emacs setup, and refreshing Go
*/
package main

import (
	"fmt"
	"strings"

	"github.com/juanman2/playgo/strutil"
)

// SayHello prints attributes bout the parameter passed in
func SayHello(foo string) {
	fmt.Printf("%s contains ll %t\n", foo, strings.Contains(foo, "hello"))
	fmt.Printf("reverse %s\n", strutil.Reverse(foo))
	fmt.Printf("%s contains ll %t\n", foo, strings.Contains(foo, "hello"))
}

func main() {
	SayHello("hello")
}
