package wordsplitter

import (
	"testing"
)

func TestReadWordFile(t *testing.T) {
	words, err := readWordFile("words.txt")
	if err != nil {
		t.Error(err)
	}
	if len(words) == 0 {
		t.Errorf("Expected at least words %d word(s) ", 1)
	}
}

func TestReverseFloatSlice(t *testing.T) {
	s1 := []float64{1, 2, 3, 4, 5}
	reversedS1 := ReverseFloatSlice(s1)
	for i, j := 0, len(s1)-1; i < len(s1); i, j = i+1, j-1 {
		if s1[i] != reversedS1[j] {
			t.Errorf("Expected %f but got %f at index %d of the reversed slice", s1[i], reversedS1[j], j)
		}
	}
}

func TestSplitRegularString(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "Howareyoudoingtoday?", want: "How are you doing today ?"},
		{input: "persistenceiskey", want: "persistence is key"},
		{input: "welcometomycity", want: "welcome to my city"},
		{input: "hundred", want: "hundred"},
	}

	for _, tc := range tests {
		result := Split(tc.input)
		if result != tc.want {
			t.Errorf("Expected %s but got %s", tc.want, result)
		}
	}
}

func TestIsInteger(t *testing.T) {
	s := "tayo12"
	isInteger(s[4])
}

func TestSplitDigitString(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "2020isthecurrentyear", want: "2020 is the current year"},
	}
	for _, tc := range tests {
		result := Split(tc.input)
		if result != tc.want {
			t.Errorf("Expected %s but got %s", tc.want, result)
		}
	}
}

func TestSplitApostropheString(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "google'sprogramminglanguage", want: "google's programming language"},
	}
	for _, tc := range tests {
		result := Split(tc.input)
		if result != tc.want {
			t.Errorf("Expected %s but got %s", tc.want, result)
		}
	}
}
