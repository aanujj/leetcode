package main

import "fmt"

func main() {

	input := "my name is anuj is again is"
	// OUTPUT should be map[again:1 anuj:1 is:3 my:1 name:1]

	countMap := make(map[string]int)
	var completeString string
	var stringBuilder = ""
	length := len(input)
	var i int
	var v rune
	for i, v = range input {

		if v != 32 {
			stringBuilder = stringBuilder + string(input[i])
		} else if i == length-1 || v == 32 {
			completeString = stringBuilder
			stringBuilder = ""
			_, found := countMap[completeString]

			if found {
				countMap[completeString]++
			} else {
				countMap[completeString] = 1
			}
		}

	}

	if stringBuilder != "" {
		countMap[stringBuilder]++
	}

	fmt.Printf("countMap: %v\n", countMap)

}
