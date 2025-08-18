package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Print("Please enter text: ")

	var s string
	fmt.Scan(&s)

	// 入力検証
	if len(s) > 100 {
		fmt.Println("No")
		return
	}

	for _, r := range s {
		if !unicode.IsLower(r) || unicode.IsDigit(r) {
			fmt.Println("No")
			return
		}
	}

	// 回文チェック
	if isPalindrome(s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}
