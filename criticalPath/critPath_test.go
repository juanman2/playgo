package critpath

import "testing"

func deepCmp(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	found := 0
	for _, ap := range a {
		for _, bp := range b {
			if (ap[0] == bp[0] && ap[1] == bp[1]) ||
				(ap[0] == bp[1] && ap[1] == bp[0]) {
				found = found + 1
				break
			}
		}
	}
	if len(a) == found {
		return true
	}

	return false
}

func TestBFCritPath(t *testing.T) {
	tests := []struct {
		inNodes       int
		inConnections [][]int
		outCritical   [][]int
	}{
		{4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}, [][]int{{1, 3}}},
		{4, [][]int{{0, 1}, {1, 2}, {2, 3}}, [][]int{{0, 1}, {1, 2}, {2, 3}}},
		{4, [][]int{{0, 1}, {0, 3}, {1, 2}, {2, 3}}, [][]int{}},
	}

	for _, tt := range tests {
		c := bfCriticalConnections(tt.inNodes, tt.inConnections)
		if deepCmp(tt.outCritical, c) == false {
			t.Errorf("bfCriticalConnections(%d , %v+) got %v+ wanted %v+",
				tt.inNodes, tt.inConnections, c, tt.outCritical)
		}
	}
}

func TestCritPath(t *testing.T) {
	tests := []struct {
		inNodes       int
		inConnections [][]int
		outCritical   [][]int
	}{
		{4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}, [][]int{{1, 3}}},
		{4, [][]int{{0, 1}, {1, 2}, {2, 3}}, [][]int{{0, 1}, {1, 2}, {2, 3}}},
		{4, [][]int{{0, 1}, {0, 3}, {1, 2}, {2, 3}}, [][]int{}},
	}

	for _, tt := range tests {
		c := criticalConnections(tt.inNodes, tt.inConnections)
		if deepCmp(tt.outCritical, c) == false {
			t.Errorf("criticalConnections(%d , %v+) got %v+ wanted %v+",
				tt.inNodes, tt.inConnections, c, tt.outCritical)
		}
	}
}
