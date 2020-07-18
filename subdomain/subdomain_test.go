// Copyright 2020 Juan Tellez All rights reserved.

package subdomain

import "testing"

// check that all the strings in slice a are also on
// slice b, regardless of order
func deepSliceCmp(a []string, b []string) bool {
	for _, ap := range a {
		found := false
		for _, bp := range b {
			if ap == bp {
				found = true
				break
			}
		}
		if found == false {
			return false
		}
	}
	return true
}

func TestSubdomainVisits(t *testing.T) {

	tests := []struct {
		in  []string
		out []string
	}{
		{
			in:  []string{"9001 discuss.leetcode.com"},
			out: []string{"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"},
		},
		{
			in:  []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
			out: []string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"},
		},
	}

	for _, tt := range tests {
		out := subdomainVisits(tt.in)
		if !deepSliceCmp(tt.out, out) {
			t.Errorf("subdomainVisits([%v]) got [%v] but expected [%v]\n",
				tt.in, out, tt.out)
		}
	}
}
