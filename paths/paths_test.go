package paths

import "testing"

func TestLengthLongestPath(t *testing.T) {

	tests := []struct {
		in  string
		out int
	}{
		{"dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext", 20},
		{"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext", 32},
	}

	for _, tt := range tests {
		if out := lengthLongestPath(tt.in); out != tt.out {
			t.Errorf("lengthLonestPath(%s) returned %d but expected %d \n",
				tt.in, out, tt.out)
		}
	}
}
