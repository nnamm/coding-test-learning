package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter text: ")

	var s string
	fmt.Scan(&s)

	if len(s) < 1 || len(s) > 1000 {
		fmt.Println("Invalid text length")
		return
	}

	var bkt []byte
	idx := 0
	for _, r := range s {
		// Open brackets
		if r == '(' || r == '[' || r == '{' {
			bkt = append(bkt, byte(r))
			idx++
		}

		// Close brackets
		if r == ')' || r == ']' || r == '}' {
			if len(bkt) > 0 {
				if bkt[idx-1] == '(' && r == ')' {
					bkt = bkt[:len(bkt)-1]
					idx--
				} else if bkt[idx-1] == '[' && r == ']' {
					bkt = bkt[:len(bkt)-1]
					idx--
				} else if bkt[idx-1] == '{' && r == '}' {
					bkt = bkt[:len(bkt)-1]
					idx--
				} else {
					break
				}
			}
		}
	}

	if len(bkt) == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
