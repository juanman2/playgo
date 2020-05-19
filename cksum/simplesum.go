package playgo

//[0,1,-3,5,8 3,4,3], 6

// Pair contains two integers which together resolve the cksum
type Pair struct {
	a int
	b int
}

// CheckSum performs a simple checksum using addition
func CheckSum(pairs []int, sum int) []Pair {

	// Allocate the lookup table
	T := make([][]bool, len(pairs))
	for i := range T {
		T[i] = make([]bool, len(pairs))
	}

	var matches []Pair

	for i := 0; i < len(pairs); i = i + 1 {
		for j := i + 1; j < len(pairs); j = j + 1 {
			if T[i][j] == true {
				continue
			}
			if sumP := pairs[i] + pairs[j]; sumP == sum {
				p := Pair{a: pairs[i], b: pairs[j]}
				matches = append(matches, p)
				T[i][j] = true
			}
		}
	}
	return matches
}

//Expected Result: [1,5], [3,3]
