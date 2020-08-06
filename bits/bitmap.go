// Copyright 2020 Juan Tellez All rights reserved.

package bitmap

// A package of functions to do bitmap operations
// these are mainly a way to remind myself of what the operations do.

import (
	"fmt"
	"unsafe"
)

func bitset(n uintptr, mem *uint) {

	var m uint
	m = 1 << (n - 1)
	*mem = *mem | m
}

// bittest checks to see if the nth bit in mem is set
func bittest(n uintptr, mem uint) (bool, error) {

	s := unsafe.Sizeof(mem)

	if n > s {
		return false, fmt.Errorf("n:%d is too big for type sizeof:%d",
			n, s)
	}

	var mask uint = 1 << (n - 1)
	if mask&mem != 0 {
		return true, nil
	}

	return false, nil
}
