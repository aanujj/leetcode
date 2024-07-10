package main

import "fmt"

func main() {

	input := "()[{}"
	output := isValid(input)
	if output {
		fmt.Println("It is valid")
	} else {
		fmt.Println("It is Not")
	}
}

//{[]}

// isValid checks if the input string has valid parentheses
func isValid(s string) bool {
	// A stack to keep track of opening brackets
	stack := []rune{}

	// A map to define matching pairs of brackets
	matchingBrackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// Iterate through each character in the string
	for _, char := range s {
		// If the character is an opening bracket, push it onto the stack
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			// If the character is a closing bracket, check for a match
			if len(stack) == 0 {
				return false
			}
			// Pop the top of the stack
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// Check if the popped bracket matches the current closing bracket
			if top != matchingBrackets[char] {
				return false
			}
		}
	}

	// If the stack is empty, all brackets matched correctly
	return len(stack) == 0
}
