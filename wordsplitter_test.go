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