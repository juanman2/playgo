package bitmap

import "testing"

func TestBittest(t *testing.T) {

	tests := []struct {
		in    uintptr
		inMem uint
		out   bool
		outE  error
	}{
		{1, 1, true, nil},
		{2, 0, false, nil},
		{3, 0xff, true, nil},
		{3, 4, true, nil},
		{3, 0x08, false, nil},
		{4, 0x08, true, nil},
	}

	for _, tt := range tests {
		if out, e := bittest(tt.in, tt.inMem); out != tt.out || e != tt.outE {
			t.Errorf("bittest(%d, %x) returned %v exptected %v",
				tt.in, tt.inMem, out, tt.out)
		}
	}

	_, e := bittest(256, 5)
	if e == nil {
		t.Errorf("bittest should have returned error but nil\n")
	}
}

func TestBitset(t *testing.T) {

	tests := []struct {
		in    uintptr
		inMem uint
		out   uint
	}{
		{1, 0, 0x01},
		{2, 0, 0x02},
		{3, 0, 0x04},
		{4, 1, 0x09},
		{1, 0x10, 0x11},
		{2, 0x0f, 0x0f},
	}

	for _, tt := range tests {
		m := tt.inMem
		bitset(tt.in, &m)
		if (tt.out & m) == 0 {
			t.Errorf("bitset(%d 0x%x) returned 0x%x expected 0x%x",
				tt.in, tt.inMem, m, tt.out)
		}
	}
}
