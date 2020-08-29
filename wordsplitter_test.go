package wordninja

import (
	"testing"
)

func TestReadWordFile(t *testing.T){
	words,err := readWordFile(wordFilePath)
	if err != nil{
		t.Error(err)
	}
	if len(words) == 0{
		t.Errorf("Expected at least words %d word(s) ",1)
	}
}

func TestReverse(t *testing.T) {
	s1 := []float64{1,2,3,4,5}
	reversedS1 := Reverse(s1)
	for i,j := 0,len(s1)-1 ; i < len(s1); i,j = i+1,j-1{
		if s1[i] != reversedS1[j]{
			t.Errorf("Expected %f but got %f at index %d of the reversed slice",s1[i],reversedS1[j],j)
		}
	}
}

func TestSplit(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "howareyoudoingtoday?", want: "how are you doing today ?"},
		{input: "persistenceiskey", want: "persistence is key"},
		{input: "hundred", want: "hundred"},
	}
	for _,tc := range  tests{
		result := Split(tc.input)
		if result != tc.want{
			t.Errorf("Expected %s but got %s",tc.want,result)
		}
	}
}