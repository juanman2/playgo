package strutil

import "testing"

func TestSentenceToCamelCap(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"hello dolly", "helloDolly"},
		{"one", "one"},
		{"", ""},
		{"BIG FAT DEAL", "bigFatDeal"},
		{"origami Gangsters", "origamiGangsters"},
		{"little endians have more fun", "littleEndiansHaveMoreFun"},
	}

	for _, tt := range tests {
		if got := SentenceToCamelCap(tt.in); got != tt.out {
			t.Errorf("SentenceToCamelCap(%v) want %v got %v",
				tt.in, tt.out, got)
		}
	}
}

func TestWordCapitalize(t *testing.T) {

	tests := []struct {
		in  string
		out string
	}{
		{"", ""},
		{"apple", "Apple"},
		{"PEACH", "Peach"},
		{"oRANGE", "Orange"},
	}

	for _, tt := range tests {
		if got := WordCapitalize(tt.in); got != tt.out {
			t.Errorf("WordCapitalize(%v) want %v got %v",
				tt.in, tt.out, got)
		}
	}
}
