package playgo

import "testing"

func pairsEquals(x []Pair, y []Pair) bool {

	if len(x) != len(y) {
		return false
	}

	foundCount := 0
	for i := 0; i < len(x); i = i + 1 {
		for j := 0; j < len(y); j = j + 1 {

			if (x[i].a == y[j].a && x[i].b == y[j].b) ||
				(x[i].a == y[j].b && x[i].b == y[j].a) {
				foundCount = foundCount + 1
				break
			}
		}
	}
	return foundCount == len(x)
}

func TestCheckSum(t *testing.T) {
	tests := []struct {
		inList  []int
		inSum   int
		outList []Pair
	}{
		{[]int{0, 1, -3, 5, 8, 3, 4, 3}, 6, []Pair{{1, 5}, {3, 3}}},
		{[]int{4, 3, 2, 6, 8, 0, 11, 1, 5}, 8, []Pair{{2, 6}, {8, 0}, {3, 5}}},
		{[]int{3, 2}, 6, []Pair{}},
	}

	for _, tt := range tests {
		pairs := CheckSum(tt.inList, tt.inSum)
		if out := pairsEquals(tt.outList, pairs); out != true {
			t.Errorf("CheckSum(%v+ , %d) got %v+ wanted %v+",
				tt.inList, tt.inSum, pairs, tt.outList)
		}

	}
}
