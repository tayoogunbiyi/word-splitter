package wordninja

import (
	"bufio"
	"math"
	"os"
)


var (
	wordFilePath string
	wordCost *CostMap
)

type CostMap struct {
	mp map[string]float64
	maxLengthKey int
}

func NewCostMap(words [] string, Cost func(int, [] string) float64) *CostMap{
	mp :=  make(map[string]float64)
	maxLengthKey := 0

	for i := 0; i < len(words); i++{
		c := Cost(i+1,words)
		mp[words[i]] = c

		if len(words[i]) > maxLengthKey{
			maxLengthKey = len(words[i])
		}
	}

	cm := &CostMap{mp,maxLengthKey}
	return cm
}


func init(){
	wordFilePath = "words.txt"
	languageWords,err := readWordFile(wordFilePath)

	if err != nil{

	}
	wordCost = NewCostMap(languageWords,func (weight int, words [] string) float64 {
		return math.Log(float64(weight)) * math.Log(float64(len(words)))
	})
}




func readWordFile(filepath string) (words [] string,err error){
	f,err := os.Open(filepath)

	if err != nil{
		return nil,err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		words = append(words,scanner.Text())
	}

	if err := scanner.Err(); err != nil{
		return nil,err
	}

	return words,nil
}