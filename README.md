# Word-splitter

This helps to split words which are joined together withoutanydelimeter.

I was working on a problem involving extracting text from some weirdly formatted PDF files then came across this really smart stack overflow answer - [How to split text without spaces into list of words ?
](https://stackoverflow.com/questions/8870261/how-to-split-text-without-spaces-into-list-of-words/11642687#11642687) which then led me to this great package - [Word Ninja](https://github.com/keredson/wordninja)

I decided to re-write it in Go.

## Installation

```
go get https://github.com/tayoogunbiyi/word-splitter
```

## Usage

```go
package main

import (
    "fmt"
	"https://github.com/tayoogunbiyi/word-splitter"
)

func main(){
    fmt.Println(wordsplitter.Split("welcometomycity")) // outputs "welcome to my city"
    fmt.Println(wordsplitter.Split("2020istheyear")) // outputs "2020 is the year"

}

```
