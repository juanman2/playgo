// Copyright 2020 Juan Tellez All rights reserved.

package maxstudents

import "testing"

func TestConsecutiveBits(t *testing.T) {

	tests := []struct {
		in  uint
		out bool
	}{
		{0x1, true},
		{0x3, false},
		{0x11, true},
		{0x13, false},
		{0x15, true},
		{0x81, true},
		{0xC0, false},
		{0x83, false},
	}

	for _, tt := range tests {
		if out := noConsecutiveBits(tt.in); out != tt.out {
			t.Errorf("noconsecutivebits(%b) got %v but expected %v\n",
				tt.in, out, tt.out)
		}
	}

}

func TestMaxStudents(t *testing.T) {

	tests := []struct {
		in  [][]byte
		out int
	}{
		{
			in: [][]byte{
				{'#', '.', '#', '#', '.', '#'},
				{'.', '#', '#', '#', '#', '.'},
				{'#', '.', '#', '#', '.', '#'},
			},
			out: 4,
		},
		{
			in: [][]byte{
				{'.', '#'},
				{'#', '#'},
				{'#', '.'},
				{'#', '#'},
				{'.', '#'},
			},
			out: 3,
		},
		{
			in: [][]byte{
				{'.', '#'},
				{'#', '#'},
				{'#', '.'},
				{'#', '#'},
				{'.', '#'},
			},
			out: 3,
		},
		{
			in: [][]byte{
				{'#', '.', '.', '.', '#'},
				{'.', '#', '.', '#', '.'},
				{'.', '.', '#', '.', '.'},
				{'.', '#', '.', '#', '.'},
				{'#', '.', '.', '.', '#'},
			},
			out: 10,
		},
		{
			in: [][]byte{
				{'#', '#', '#', '.', '#'},
				{'.', '.', '#', '.', '.'},
				{'#', '.', '#', '.', '#'},
				{'.', '.', '.', '.', '.'},
				{'.', '.', '.', '#', '.'},
			},
			out: 9,
		},
		{
			in: [][]byte{
				{'#', '.', '.'},
				{'#', '.', '.'},
				{'#', '#', '.'},
			},
			out: 3,
		},
		{
			in: [][]byte{
				{'.', '#', '#', '#', '#', '.', '.'},
				{'#', '.', '.', '.', '#', '.', '#'},
				{'.', '.', '#', '.', '#', '.', '#'},
				{'#', '#', '#', '#', '.', '.', '#'},
				{'#', '#', '.', '#', '#', '#', '#'},
				{'.', '#', '#', '#', '#', '.', '.'},
			},
			out: 11,
		},
		{
			in: [][]byte{
				{'#', '.', '.'},
				{'#', '.', '.'},
				{'.', '#', '#'},
			},
			out: 3,
		},
		{
			in: [][]byte{
				{'.', '.', '.', '.', '#', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '#', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '#', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '#', '.', '.', '#', '.'},
			},
			out: 31,
		},
	}

	// DP approach using bitmasks
	for _, tt := range tests {
		if got := maxStudents(tt.in); got != tt.out {
			t.Errorf("maxStudents(%q) got %d but expected %d \n",
				tt.in, got, tt.out)
		}
	}

	// Heuristic approach
	for _, tt := range tests {
		if got := maxStudentsHeuristic(tt.in); got != tt.out {
			t.Errorf("maxStudents(%q) got %d but expected %d \n",
				tt.in, got, tt.out)
		}
	}

}
