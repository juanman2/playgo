// Copyright 2020 Juan Tellez All rights reserved.

package maxstudents

import (
	"math/bits"
)

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

Constraints:

seats contains only characters '.' and'#'.

o m == seats.length
o n == seats[i].length
o 1 <= m <= 8
o 1 <= n <= 8

Solution:

The solution below is a DP brute force bitmap scheduling solution.
It attempts all combinations on a bitmap and uses efficient masking
to test the solutions.
*/

// The test constraint is 8x8 so we use uint everywhere
type seatSchedule struct {
	rowMax int
	colMax int
	seats  [][]byte

	maxSeatings  uint
	badSeatsMask []uint
	rowMask      []uint
	memo         []map[uint]int
}

// noConsecutiveBits makes sure a seat has no neighbors
// it does this by taking copy of the mask and shift it left
// then taking a copy of the mask and shift it right
// bitwise and the mask with both copies, if either
// comes out true, then there were two consecutive
// bits somewhere.  We go both ways to deal with the edge
// bits.
func noConsecutiveBits(r uint) bool {
	tmp1 := r
	tmp1 = tmp1 << 1
	tmp2 := r
	tmp2 = tmp2 >> 1

	if (r&tmp1) != 0 ||
		(r&tmp2) != 0 {
		return false
	}
	return true
}

func (s seatSchedule) validateSeating(rowMask uint, prevRowMask uint) bool {

	// check row behind
	if ((prevRowMask<<1)&rowMask) == 0 &&
		((prevRowMask>>1)&rowMask) == 0 {
		return true
	}
	return false
}

func (s seatSchedule) scheduleRow(row uint, prevRowMask uint) int {

	maxCount := int(0)

	// if another recursion tree has already computed this, return it.
	if c, ok := s.memo[row][prevRowMask]; ok {
		return c
	}
	maxCount = 0
	for smask := uint(0); smask < s.maxSeatings; smask = smask + 1 {

		// if this arrangent tries to seat anyone on a bad seat, skip it
		if (smask & s.badSeatsMask[row]) != 0 {
			continue
		}

		// validate that the seating has no neighbors to left and right
		if noConsecutiveBits(smask) == false {
			continue
		}

		//	Validate against the previous row
		if prevRowMask != 0 && s.validateSeating(smask, prevRowMask) == false {
			continue
		}

		// This is a viable seating, so save the count
		count := bits.OnesCount(smask)

		// recurse
		if int(row+1) < s.rowMax {
			count = s.scheduleRow(row+1, smask) + count
		}

		if count > maxCount {
			s.rowMask[row] = smask
			maxCount = count
		}
	}

	s.memo[row][prevRowMask] = maxCount
	// fmt.Printf("Returning %d - [%*.*b] [%*.*b] row:%d\n",
	// 	maxCount, s.colMax, s.colMax, prevRowMask,
	// 	s.colMax, s.colMax, s.rowMask[row], row)

	return maxCount
}

func maxStudents(seats [][]byte) int {

	//	fmt.Printf("maxStudents this:%q\n", seats)

	s := new(seatSchedule)
	s.seats = seats

	s.rowMax = len(seats)
	s.colMax = len(seats[0])

	// The test constraint is 8x8
	s.badSeatsMask = make([]uint, s.rowMax)

	// make the bad seats mask
	for r, row := range seats {
		for c, seat := range row {
			if seat == '#' {
				// set the bit
				m := uint(1) << c
				s.badSeatsMask[r] = s.badSeatsMask[r] | m
			}
		}
	}
	//	fmt.Printf("Bad Seats Mask: %*.*b\n", s.colMax, s.colMax, s.badSeatsMask)

	// loop over total number of possible combinations per row
	// Maximum number of seat combinations
	s.maxSeatings = (uint(1) << s.colMax)
	s.memo = make([]map[uint]int, s.rowMax)
	for i := range s.memo {
		s.memo[i] = make(map[uint]int)
	}
	s.rowMask = make([]uint, s.rowMax)

	count := s.scheduleRow(0, 0)

	//	fmt.Printf("For %q Returning:%d\n", seats, count)

	return count
}
