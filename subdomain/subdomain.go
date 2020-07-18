// Copyright 2020 Juan Tellez All rights reserved.

package subdomain

import (
	"fmt"
	"strconv"
)

/*

  A website domain like "discuss.leetcode.com" consists of various
  subdomains. At the top level, we have "com", at the next level, we
  have "leetcode.com", and at the lowest level,
  "discuss.leetcode.com". When we visit a domain like
  "discuss.leetcode.com", we will also visit the parent domains
  "leetcode.com" and "com" implicitly.

  Now, call a "count-paired domain" to be a count (representing the
  number of visits this domain received), followed by a space, followed
  by the address. An example of a count-paired domain might be "9001
  discuss.leetcode.com".

  We are given a list cpdomains of count-paired domains. We would like a
  list of count-paired domains, (in the same format as the input, and in
  any order), that explicitly counts the number of visits to each
  subdomain.
*/

// stringTokenize returns an array of indexes to a string where
// the token t is found
func stringTokenize(t1 rune, t2 rune, str string) []int {
	idxs := make([]int, 0, 3)

	for i, c := range str {
		if c == t1 || c == t2 {
			idxs = append(idxs, i)
		}
	}

	return idxs
}

func subdomainVisits(cpdomains []string) []string {

	domains := make(map[string]int, 0)

	if len(cpdomains) == 0 {
		return []string{}
	}

	for _, input := range cpdomains {

		var s string
		var countStr string
		idxs := stringTokenize('.', ' ', input)
		countStr = input[:idxs[0]]
		count, err := strconv.Atoi(countStr)
		if err != nil {
			panic(err)
		}
		s = input[idxs[0]+1:]

		domains[s] = domains[s] + count
		for _, i := range idxs[1:] {
			i++
			domains[input[i:]] = domains[input[i:]] + count
		}
	}

	out := make([]string, 0)
	for d, c := range domains {
		s := fmt.Sprintf("%d %s", c, d)
		out = append(out, s)
	}

	return out
}
