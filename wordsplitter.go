package wordsplitter

import (
	"bufio"
	"math"
	"os"
	"strings"
	"fmt"
)


var (
	wordFilePath string
	wordCost *CostMap
)

type CostMap struct {
	mp map[string]float64
	maxLengthKey int
}

func (cm *CostMap) Get (key string,default_ float64) float64{
	val, ok := cm.mp[key]
	if !ok {
		return default_
	}
	return val
}

func init(){
	wordFilePath = "words.txt"
	languageWords,err := readWordFile(wordFilePath)

	if err != nil {
		fmt.Println("Could not load word frequency list - " + err.Error())
		os.Exit(1)
	}
	wordCost = NewCostMap(languageWords,func (weight int, words [] string) float64 {
		return math.Log(float64(weight)) * math.Log(float64(len(words)))
	})
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

func ReverseFloatSlice(s [] float64) [] float64 {
	reversedSlice := make([]float64,0,len(s))
	j := len(s)-1

	for j >= 0{
		reversedSlice = append(reversedSlice,s[j])
		j-=1
	}
	return reversedSlice
}

func ReverseStringSlice(s [] string) [] string {
	reversedSlice := make([]string,0,len(s))
	j := len(s)-1

	for j >= 0{
		reversedSlice = append(reversedSlice,s[j])
		j-=1
	}
	return reversedSlice
}

func FindBestMatch(s string,costIdx int, cost [] float64) (float64,int){
	startIdx := int(math.Max(0,float64(costIdx-wordCost.maxLengthKey)))

	candidatesForBestMatch := cost[startIdx:costIdx]
	candidatesForBestMatch = ReverseFloatSlice(candidatesForBestMatch)
	optimalCost := math.Inf(1)
	optimalCostIdx := -1
	for k,c := range candidatesForBestMatch{
		key := strings.ToLower(s[costIdx-k-1:costIdx])
		newCost := c + wordCost.Get(key,9.0e99)
		if newCost < optimalCost{
			optimalCost = newCost
			optimalCostIdx = k+1
		}
	}

	return optimalCost,optimalCostIdx
}

func isInteger(ch uint8) bool{
	return ch >= 48 && ch <= 57
}

func Split(s string) string{
	cost := []float64 {0.0}
	for i := 1; i < len(s)+1; i++{
		c,_ := FindBestMatch(s,i,cost)
		cost = append(cost,c)
	}
	var output []string
	for i := len(s); i > 0;{
		_,k := FindBestMatch(s,i,cost)
		currSubstring := s[i-k:i]
		nT := true
		if currSubstring != "'"{
			if len(output) > 0{
				if output[len(output)-1] == "'s" || (isInteger(s[i-1]) && isInteger(output[len(output)-1][0])){
					output[len(output)-1] = currSubstring + output[len(output)-1]
					nT = false
				}
			}
		}
		if nT{
			output = append(output, currSubstring)
		}
		i-=k
	}
	return strings.Join(ReverseStringSlice(output)," ")


}