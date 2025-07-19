package main

import "fmt"

func main() {
	fmt.Print("Enter text: ")

	var s string
	fmt.Scan(&s)

	if isValidBrackets(s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func isValidBrackets(s string) bool {
	var stack []rune
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, r := range s {
		// Open brackets
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r) // Push
		}

		// Close brackets
		if r == ')' || r == ']' || r == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[r] {
				return false
			}
			stack = stack[:len(stack)-1] // Pop
		}
	}

	return len(stack) == 0
}
