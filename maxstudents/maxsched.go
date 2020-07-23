// Copyright 2020 Juan Tellez All rights reserved.

package maxstudents

import "fmt"

/*

Given a m * n matrix seats that represent seats distributions in a
classroom. If a seat is broken, it is denoted by '#' character
otherwise it is denoted by a '.' character.

Students can see the answers of those sitting next to the left, right,
upper left and upper right, but he cannot see the answers of the
student sitting directly in front or behind him. Return the maximum
number of students that can take the exam together without any
cheating being possible..

Students must be placed in seats in good condition.

Solution:

The solution below is a heuristic solution, using different scheduling
algorithms.  On Leetcode it scored highly but I suspect there are not
many submissions for this problem.  Please see the DP solution.

*/

// schedulebyColumn has four modes, start with 0 column, 1 column
// try to schedule on all columns, column by column, or skip all odds
// skip all evens
func scheduleByColumn(s [][]byte, odd bool, simple bool) int {
	maxCol := len(s[0])

	maxRow := len(s)

	t := 0

	start := 0
	if odd == true {
		start = 1
	}

	i := 1
	if simple == true {
		i = 2
	}

	for c := start; c < maxCol; c = c + i {
		for r := 0; r < maxRow; r = r + 1 {
			if validSeat(s, r, c) {
				s[r][c] = 'x'
				t++
			}
		}
	}

	return t
}

func scheduleByColumnRightToLeft(s [][]byte) int {
	maxCol := len(s[0])
	maxRow := len(s)

	t := 0

	for c := maxCol - 1; c >= 0; c = c - 1 {
		for r := 0; r < maxRow; r = r + 1 {
			if validSeat(s, r, c) {
				s[r][c] = 'x'
				t++
			}
		}
	}

	return t
}

func scheduleByRow(s [][]byte) int {
	maxCol := len(s[0])
	maxRow := len(s)

	t := 0

	for r := 0; r < maxRow; r = r + 1 {
		for c := 0; c < maxCol; c = c + 1 {
			if validSeat(s, r, c) {
				s[r][c] = 'x'
				t++
			}
		}
	}

	return t
}

func scheduleByRowBottomUp(s [][]byte) int {
	maxCol := len(s[0])
	maxRow := len(s)

	t := 0
	for r := maxRow - 1; r >= 0; r = r - 1 {
		for c := 0; c < maxCol; c = c + 1 {
			if validSeat(s, r, c) {
				s[r][c] = 'x'
				t++
			}
		}
	}

	return t
}

// schedulebyColumnOrdered schedules the columns that have the most free legal seats first.
func scheduleByColumnOrdered(s [][]byte) int {
	countMap := make(map[int][]int)
	maxCol := len(s[0])
	maxRow := len(s)

	for c := 0; c < maxCol; c = c + 1 {
		count := 0
		for r := 0; r < maxRow; r = r + 1 {
			if s[r][c] == '.' {
				count++
			}
		}
		if countMap[count] == nil {
			countMap[count] = make([]int, 0)
		}
		countMap[count] = append(countMap[count], c)
	}

	t := 0
	for count := maxRow; count >= 0; count = count - 1 {
		if countMap[count] != nil {
			for _, c := range countMap[count] {
				for r := 0; r < maxRow; r = r + 1 {
					if validSeat(s, r, c) {
						s[r][c] = 'x'
						t++
					}
				}
			}
		}
	}
	return t
}

// neighborCount returns the number of legal seats next to the current seat
func neighborCount(s [][]byte, row int, col int) int {

	// if not a validseat put it in max seat map slot
	if s[row][col] != '.' {
		return 7
	}

	maxRow := len(s) - 1
	maxCol := len(s[0]) - 1
	n := 0

	if col < maxCol && s[row][col+1] == '.' { // neighbor to right
		n++
	}
	if col < maxCol && row > 0 &&
		s[row-1][col+1] == '.' { // neighbor to upper right
		n++
	}
	if col < maxCol && row < maxRow &&
		s[row+1][col+1] == '.' { // neighbor to lower right
		n++
	}
	if col > 0 && row > 0 &&
		s[row-1][col-1] == '.' { // neighbor to the upper left
		n++
	}
	if col > 0 && s[row][col-1] == '.' { // neighbor to the left
		n++
	}

	if col > 0 && row < maxRow &&
		s[row+1][col-1] == '.' { // neighbor to the lower left
		n++
	}

	return n
}

type rcPair struct {
	row int
	col int
}

// scheduleByNeighborCount schedules the students to seats that have the smallest number
// of legal neighbors first.
func scheduleByNeighborCount(s [][]byte) int {
	countMap := make(map[int][]rcPair)
	maxCol := len(s[0])
	maxRow := len(s)

	for c := 0; c < maxCol; c = c + 1 {
		count := 0
		for r := 0; r < maxRow; r = r + 1 {
			count = neighborCount(s, r, c)
			rc := rcPair{r, c}
			if countMap[count] == nil {
				countMap[count] = make([]rcPair, 0)
			}
			countMap[count] = append(countMap[count], rc)
		}
	}

	t := 0
	for count := 0; count < 7; count = count + 1 {
		if countMap[count] != nil {
			for _, c := range countMap[count] {
				for r := 0; r < maxRow; r = r + 1 {
					if validSeat(s, c.row, c.col) {
						s[c.row][c.col] = 'x'
						t++
					}
				}
			}
		}
	}
	return t
}

func validSeat(s [][]byte, row int, col int) bool {

	if s[row][col] != '.' {
		return false
	}

	maxRow := len(s) - 1
	maxCol := len(s[0]) - 1

	switch {
	case col < maxCol && s[row][col+1] == 'x': // neighbor to right
		return false
	case col < maxCol && row > 0 &&
		s[row-1][col+1] == 'x': // neighbor to upper right
		return false
	case col < maxCol && row < maxRow &&
		s[row+1][col+1] == 'x': // neighbor to lower right
		return false
	case col > 0 && row > 0 &&
		s[row-1][col-1] == 'x': // neighbor to the upper left
		return false
	case col > 0 && s[row][col-1] == 'x': // neighbor to the left
		return false
	case col > 0 && row < maxRow &&
		s[row+1][col-1] == 'x': // neighbor to the lower left
		return false
	}
	//	fmt.Printf("[%d,%d]", row, col)
	return true
}

func cleanup(s [][]byte) {
	for r := 0; r < len(s); r = r + 1 {
		for c := 0; c < len(s[0]); c = c + 1 {
			if s[r][c] == 'x' {
				s[r][c] = '.'
			}
		}
	}
}

func maxStudentsHeuristic(seats [][]byte) int {

	max := 0

	got := scheduleByColumn(seats, true, false)
	//	fmt.Printf("For %q ... byCol odd got %d\n", seats, got)
	cleanup(seats)
	if got > max {
		max = got
	}

	got = scheduleByColumn(seats, true, true)
	// fmt.Printf("For %q ... byCol odd simple got %d\n", seats, got)
	cleanup(seats)
	if got > max {
		max = got
	}

	got = scheduleByColumn(seats, false, false)
	// fmt.Printf("For %q ... byCol even got %d\n", seats, got)
	cleanup(seats)
	if got > max {
		max = got
	}

	got = scheduleByColumn(seats, false, true)
	// fmt.Printf("For %q ... byCol even simple got %d\n", seats, got)
	cleanup(seats)
	if got > max {
		max = got
	}

	got = scheduleByRow(seats)
	// fmt.Printf("For %q ... byRow got %d\n", seats, got)
	cleanup(seats)
	if got > max {
		max = got
	}

	got = scheduleByColumnRightToLeft(seats)
	// fmt.Printf("For %q ... byCol RtoL got %d\n", seats, got)
	cleanup(seats)
	if got > max {
		max = got
	}

	got = scheduleByRowBottomUp(seats)
	// fmt.Printf("For %q ... byRow BtmUp got %d\n", seats, got)
	cleanup(seats)
	if got > max {
		max = got
	}

	got = scheduleByColumnOrdered(seats)
	// fmt.Printf("For %q ... byCol Ordered got %d\n", seats, got)
	cleanup(seats)
	if got > max {
		max = got
	}

	got = scheduleByNeighborCount(seats)
	// fmt.Printf("For %q ... byNeighbor Count got %d\n", seats, got)
	cleanup(seats)
	if got > max {
		max = got
	}

	fmt.Printf("Returning %d\n", max)

	return max
}
