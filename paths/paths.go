package paths

const (
	readFileState  = 0
	readTokenState = 1
	genPathState   = 2
)

type stack struct {
	names []int
	level int
}

func (s *stack) push(flen int) {
	if s.level < len(s.names) {
		s.names[s.level] = flen
	} else {
		s.names = append(s.names, flen)
	}
	s.level++
}

func (s *stack) pop() {
	s.level--
	s.names[s.level] = 0
}

// calculates the max len of the path
func (s *stack) genPath() int {
	pathlen := 0
	for l := 0; l < s.level; l++ {
		if l != 0 {
			pathlen = pathlen + 1
		}
		pathlen = pathlen + s.names[l]
	}

	return pathlen
}

func lengthLongestPath(input string) int {

	stk := new(stack)
	stk.names = make([]int, 0)
	stk.level = 0
	idx, start, maxLen := 0, 0, 0
	state := readFileState

	readFile := func() int {
		newState := readTokenState
		for start = idx; idx < len(input) && input[idx] != '\n'; idx++ {
			if input[idx] == '.' {
				newState = genPathState
			}
		}
		stk.push(idx - start)
		return newState
	}

	readToken := func() int {
		level := 0
		if input[idx] == '\n' {
			idx++
		}
		for ; input[idx] == '\t'; idx++ {
			level++
		}
		for level < stk.level {
			stk.pop()
		}
		return readFileState
	}

	for idx < len(input) || state == genPathState {

		switch state {
		case readFileState:
			state = readFile()
		case readTokenState:
			state = readToken()
		case genPathState:
			pathLen := stk.genPath()
			if pathLen > maxLen {
				maxLen = pathLen
			}
			state = readTokenState
		}
	}
	return maxLen
}
