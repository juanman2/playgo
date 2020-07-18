// Copyright 2020 Juan Tellez All rights reserved.

package lrucache

import "testing"

func TestLRUCache(t *testing.T) {
	tests := []struct {
		op  string
		in  []int
		out int
	}{
		{"INI", []int{2}, 0},
		{"PUT", []int{1, 1}, 0},
		{"PUT", []int{2, 2}, 0},
		{"GET", []int{1}, 1},
		{"PUT", []int{3, 3}, 0}, // evicts 2
		{"GET", []int{2}, -1},
		{"PUT", []int{4, 4}, 0}, // evicts 1
		{"GET", []int{1}, -1},
		{"GET", []int{3}, 3},
		{"GET", []int{4}, 4},
	}

	var c LRUCache

	for _, tt := range tests {
		op := tt.op
		switch op {
		case "INI":
			c = Constructor(tt.in[0])
		case "PUT":
			c.Put(tt.in[0], tt.in[1])
		case "GET":
			got := c.Get(tt.in[0])
			if got != tt.out {
				t.Errorf("LRUCache.Get failed(%v) returned %v expected %v", tt.in[0], got, tt.out)
			}
		}
	}
}
