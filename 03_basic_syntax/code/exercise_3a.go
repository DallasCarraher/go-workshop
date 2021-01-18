package main

import "fmt"

func main() {
	mySentence := "Hello, my name is Dallas"
	for index, sentence := range mySentence {
		if index%2 == 1 {
			fmt.Print(string(sentence))
		}
	}
}
